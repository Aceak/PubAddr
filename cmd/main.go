package main

import (
	"PubAddr/internal/config"
	"PubAddr/internal/logger"
	"PubAddr/internal/server"
	"os/signal"
	"syscall"

	"context"
	"net/http"
	"time"
)

func main() {
	logger.InitLogger("info")

	cfg, err := config.Load("./config.yaml")
	if err != nil {
		logger.Fatal("Failed to load config: %v", err)
	}

	logger.SetLevel(cfg.Server.LogLevel)

	srv, err := server.NewHTTPServer(cfg)

	if err != nil {
		logger.Fatal("Failed to create HTTP server: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done()
	logger.Debug("Received shutdown signal...")

	// 给 Shutdown 预留 5 秒超时
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("Graceful shutdown failed: %v", err)
	} else {
		logger.Debug("HTTP server gracefully stopped.")
	}
}
