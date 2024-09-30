package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kevinharv/vgrp-stack/server/utils"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // To-Do - get log level from environment
	})
	logger := slog.New(handler)
	logger.Info("Starting server")

	config := utils.GetConfiguration(logger)

	var tlsConfig *tls.Config
	if config.TLSEnabled {
		logger.Info("Loading TLS keypair")
		
		cert, err := tls.LoadX509KeyPair(config.TLSCertPath, config.TLSKeyPath)
		if err != nil {
			logger.Error("Failed to load TLS certificate")
		}

		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
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

	mux := http.NewServeMux()

	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", "", config.ServerPort),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig,
	}

	// Create channel and handle OS signals
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)


	// Run HTTP server in Goroutine
	go func(l *slog.Logger) {
		err := s.ListenAndServeTLS(config.TLSCertPath, config.TLSKeyPath)
		if err != nil && err != http.ErrServerClosed {
			l.Error("HTTP server closed unexpectedly\n")
			l.Error(fmt.Sprintf("%s\n", err))
		}
	}(logger)

	// Log ONLY when the server has started
	logger.Info("Started HTTP server")

	// Shutdown the server on OS interrupts/calls
	<-exitChannel
	logger.Info("Stopping server")

	// Create context and give 5 seconds to timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	// Shutdown the HTTP server
	if err := s.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", "HTTP server error", err)
	}
	logger.Info("Goodbye")
}
