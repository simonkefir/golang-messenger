package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_server "github.com/simonkefir/golang-messenger/internal/core/transport/http/server"
	chats_repository_postgres "github.com/simonkefir/golang-messenger/internal/feature/chats/repository/postgres"
	chats_service "github.com/simonkefir/golang-messenger/internal/feature/chats/service"
	chats_transport_http "github.com/simonkefir/golang-messenger/internal/feature/chats/transport/http"
	messages_repository_postgres "github.com/simonkefir/golang-messenger/internal/feature/messages/repository/postgres"
	messages_service "github.com/simonkefir/golang-messenger/internal/feature/messages/service"
	messages_transport_http "github.com/simonkefir/golang-messenger/internal/feature/messages/transport/http"
	users_repository_postgres "github.com/simonkefir/golang-messenger/internal/feature/users/repository/postgres"
	users_service "github.com/simonkefir/golang-messenger/internal/feature/users/service"
	users_transport_http "github.com/simonkefir/golang-messenger/internal/feature/users/transport/http"
	"go.uber.org/zap"
)

const minJWTSecretLength = 32

func main() {
	setTimezone()
	validateJWTEnv()

	logCfg, err := core_logger.NewConfig()
	if err != nil {
		log.Fatalf("logger config: %v", err)
	}

	logger, err := core_logger.NewLogger(logCfg)
	if err != nil {
		log.Fatalf("logger: %v", err)
	}
	defer logger.Close()

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("db open", zap.Error(err))
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Fatal("db ping", zap.Error(err))
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	userRepo := users_repository_postgres.NewUserRepository(db)
	userService := users_service.NewUsersService(userRepo)
	userHandler := users_transport_http.NewUsersHTTPHandler(userService)

	chatRepo := chats_repository_postgres.NewChatRepository(db)
	chatService := chats_service.NewChatsService(chatRepo)
	chatHandler := chats_transport_http.NewChatsHTTPHandler(chatService)

	msgRepo := messages_repository_postgres.NewMsgRepository(db)
	msgService := messages_service.NewMessagesService(msgRepo, chatRepo)
	msgHandler := messages_transport_http.NewMessagesHTTPHandler(msgService)

	v1 := core_http_server.NewAPIVersionRouter(
		core_http_server.ApiVersion1,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	v1.RegisterRoutes(userHandler.Routes()...)

	v1.RegisterRoutes(chatHandler.Routes()...)

	v1.RegisterRoutes(msgHandler.Routes()...)

	cfg, err := core_http_server.NewConfig()
	if err != nil {
		logger.Fatal("server config", zap.Error(err))
	}

	srv := core_http_server.NewHTTPServer(cfg)
	srv.RegisterAPIRouters(v1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		logger.Info("shutting down...")
		cancel()
	}()

	if err := srv.Run(ctx); err != nil {
		logger.Fatal("server", zap.Error(err))
	}
}

func setTimezone() {
	tz := os.Getenv("TZ")
	if tz == "" {
		return
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatalf("invalid TZ: %v", err)
	}
	time.Local = loc
}

func validateJWTEnv() {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) < minJWTSecretLength {
		log.Fatalf("JWT_SECRET must be at least %d characters, got %d", minJWTSecretLength, len(secret))
	}

	ttl := os.Getenv("JWT_TTL")
	if ttl == "" {
		log.Fatal("JWT_TTL is not set")
	}

	if _, err := time.ParseDuration(ttl); err != nil {
		log.Fatalf("invalid JWT_TTL: %v", err)
	}
}
