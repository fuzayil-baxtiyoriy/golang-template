package main

import (
	"context"
	"errors"
	"golang-template/internal/config"
	httpsrv "golang-template/internal/handler/http"
	"golang-template/internal/infra/postgres"
	"golang-template/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// load config files
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	logger, err := logger.NewLogger(cfg.Logger.Level, cfg.Logger.Format)

	if err != nil {
		log.Fatalf("failed to create logger instance, %v", err)
	}
	logger.Info("Initialized logger instance")

	// Initialize pgxpool.Pool
	pool, err := postgres.NewConnPool(cfg)
	if err != nil {
		log.Fatalf("failed to initialize connection pool, %v", err)
	}
	// just remove when adding service layer
	defer pool.Close()
	logger.Info("Initialized connection pool")

	// Build DSN from config
	dsn := postgres.BuildDSN(cfg)

	// Apply DB migrations
	migrationPath := "file://migrations" // adjust path if needed
	err = postgres.MigratePostgresDB(migrationPath, dsn)
	if err != nil {
		log.Fatalf("failed to run DB migrations: %v", err)
	}
	logger.Info("DB migrations applied")

	// Start HTTP server
	httpSrvAddr := "localhost:8080"
	httpSrv, err := httpsrv.New(cfg, logger, httpSrvAddr)
	if err != nil {
		logger.Fatalf("failed to start HTTP server: %v", err)
	}
	go func() {
		if err := httpSrv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("error occurred while running http server: %w", err)
		}
	}()
	logger.Info("HTTP server started at address: " + httpSrvAddr)

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		logger.Fatalf("HTTP server shutdown failed: %v", err)
	}
	logger.Info("HTTP server exited properly")
}
