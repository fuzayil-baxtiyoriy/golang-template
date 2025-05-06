package main

import (
	"golang-template/internal/config"
	"golang-template/internal/infra/postgres"
	"golang-template/pkg/logger"
	"log"
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
}
