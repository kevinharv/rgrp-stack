package main

import (
	"log/slog"
	"os"

	"github.com/kevinharv/vgrp-stack/server/utils"
)


func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	
	logger.Info("Starting server")

	config := utils.GetConfiguration()

	if config.TLSEnabled {
		logger.Info("Loading TLS keypair")
	}

	// Get configuration from environment
	// Setup TLS if enabled
	// Connect to Postgres DB
	// Connect to Redis

	// TODO - populate Redis cache

	// Setup routes

	// Start server

	// Listen for stop signals - close gracefully
		// Flush Redis to Postgres (if needed)
		// Close Redis connection
		// Close Postgres connection
		// Shutdown HTTP server
}