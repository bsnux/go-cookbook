package main

import (
	"log/slog"
	"os"
)

func main() {
	// Using default logger where `Info` is the default level. This means
	// `slog.Debug` won't print anything
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

	// Creating a new `logfmt` standard logger and setting default level to debug
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger.Debug("debug info")
	logger.Info("test info")

	// Creating new JSON logger and setting default level to warning
	logJ := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	logJ.Info("info msg")
	logJ.Warn("warning msg")
}
