package core_http_server

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
}

func NewConfig() (Config, error) {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		return Config{}, fmt.Errorf("HTTP_ADDR is required")
	}

	shutdownTimeout := 30 * time.Second
	if v := os.Getenv("HTTP_SHUTDOWN_TIMEOUT"); v != "" {
		duration, err := time.ParseDuration(v)
		if err != nil {
			return Config{}, fmt.Errorf("invalid HTTP_SHUTDOWN_TIMEOUT: %w", err)
		}
		shutdownTimeout = duration
	}

	return Config{
		Addr:            addr,
		ShutdownTimeout: shutdownTimeout,
	}, nil
}
