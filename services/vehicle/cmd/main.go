package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/config"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/logger"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/server"
)

// main is the entry point of the vehicle service
func main() {

	// Prepare server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	logger.Init()
	cfg := config.New()
	srv, err := server.New(ctx, cfg)
	if err != nil {
		logger.Log.Error("failed to create server", slog.String("error", err.Error()))
		return
	}

	// Start server in a separate goroutine
	go func() {
		if err := srv.Start(); err != nil {
			logger.Log.Error("failed to start server", slog.String("error", err.Error()))
			stop()
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-ctx.Done()
	logger.Log.Info("shutting down server")
	if err := srv.Stop(); err != nil {
		logger.Log.Error("failed to gracefully stop server", slog.String("error", err.Error()))
	}
}
