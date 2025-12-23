package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/config"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/logger"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/server"
)

// main initializes and starts the Inspection Service gRPC/HTTP server
func main() {

	// Prepare server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	logger.Init()
	cfg := config.New()
	srv, err := server.New(cfg)
	if err != nil {
		logger.Log.Error("failed to create server", slog.String("error", err.Error()))
		return
	}

	// Start server in a separate goroutine
	go func() {
		if err := srv.Start(ctx); err != nil {
			logger.Log.Error("failed to start server", slog.String("error", err.Error()))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-ctx.Done()
	if err := srv.Stop(); err != nil {
		logger.Log.Error("failed to stop server", slog.String("error", err.Error()))
	}
}
