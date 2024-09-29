package main

import (
	"log/slog"
	"os"
)


func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)

	logger.Info("Started Server")
}