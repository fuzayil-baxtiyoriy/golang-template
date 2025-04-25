package main

import (
	"golang-template/pkg/logger"
	"log"
)

func main() {
	logger, err := logger.NewLogger("debug", "text")

	if err != nil {
		log.Fatalf("failed to create logger instance, %v", err)
	}
	logger.Info("Initialized logger instance")
}
