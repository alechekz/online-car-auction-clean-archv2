package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/config"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/logger"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/server"
)

// main initializes and starts the Inspection Service HTTP server
func main() {

	// Prepare server
	logger.Init()
	cfg := config.New()
	srv, err := server.New(cfg)
	if err != nil {
		logger.Log.Error("failed to create server", slog.String("error", err.Error()))
		return
	}

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a separate goroutine
	go func() {
		if err := srv.Start(); err != nil {
			logger.Log.Error("failed to start server", slog.String("error", err.Error()))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-stop
	if err := srv.Stop(); err != nil {
		logger.Log.Error("failed to stop server", slog.String("error", err.Error()))
	}

}
