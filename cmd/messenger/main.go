package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	core_http_server "github.com/simonkefir/golang-messenger/internal/core/transport/http/server"
	users_repository_postgres "github.com/simonkefir/golang-messenger/internal/feature/users/repository/postgres"
	users_service "github.com/simonkefir/golang-messenger/internal/feature/users/service"
	users_transport_http "github.com/simonkefir/golang-messenger/internal/feature/users/transport/http"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("db ping: %v", err)
	}

	userRepo := users_repository_postgres.NewUserRepository(db)
	userService := users_service.NewUsersService(userRepo)
	userHandler := users_transport_http.NewUsersHTTPHandler(userService)

	v1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	v1.RegisterRoutes(userHandler.Routes()...)

	cfg, err := core_http_server.NewConfig()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	srv := core_http_server.NewHTTPServer(cfg)
	srv.RegisterAPIRouters(v1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		cancel()
	}()

	if err := srv.Run(ctx); err != nil {
		log.Fatalf("server: %v", err)
	}
}
