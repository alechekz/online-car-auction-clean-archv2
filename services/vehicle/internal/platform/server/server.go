package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/config"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/logger"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/postgres"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/provider/inspectionclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/provider/pricingclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/service"
	vehicleHttp "github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/transport/http"
)

// Server represents the HTTP server
type Server struct {
	httpServer *http.Server
	ctx        context.Context
}

// New creates a new Server instance
func New(ctx context.Context, cfg *config.Config) (*Server, error) {

	//open database connection
	repo, err := postgres.NewVehicleRepo(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Log.Error("failed to connect to postgres", slog.String("error", err.Error()))
		return nil, err
	}

	//init inspection gRPC client
	inspection, err := inspectionclient.New(cfg.InspectionURL)
	if err != nil {
		logger.Log.Error("failed to create inspection gRPC client", slog.String("error", err.Error()))
		return nil, err
	}
	// defer inspection.Close() //nolint:errcheck

	//init pricing gRPC client
	pricing, err := pricingclient.New(cfg.PricingURL)
	if err != nil {
		logger.Log.Error("failed to create pricing gRPC client", slog.String("error", err.Error()))
		return nil, err
	}
	// defer pricing.Close() //nolint:errcheck

	//init usecases
	uc := service.NewVehicleUC(repo, inspection, pricing)
	bulkUc := service.NewVehiclesBulkUC(repo, uc)

	//init Gin engine
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	vehicleHttp.RegisterVehicleRoutes(r, uc, bulkUc)

	//create HTTP server
	srv := &http.Server{
		Addr:              cfg.Address,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
	return &Server{ctx: ctx, httpServer: srv}, nil
}

// Start runs the HTTP server on the specified address
func (s *Server) Start() error {
	logger.Log.Info("starting Gin server", slog.String("address", s.httpServer.Addr))
	return s.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the HTTP server
func (s *Server) Stop() error {
	logger.Log.Info("stopping Gin server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		return err
	}
	return nil
}

// Handler returns the HTTP handler for testing purposes
func (s *Server) Handler() http.Handler {
	return s.httpServer.Handler
}
