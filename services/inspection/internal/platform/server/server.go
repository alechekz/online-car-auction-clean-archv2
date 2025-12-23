package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	grpcDelivery "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/transport/gRPC"
	pb "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/transport/gRPC/proto"

	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/config"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/logger"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/provider/builddataclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/provider/msrpdataclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/service"
)

// Server represents both HTTP and gRPC servers for the Inspection Service
type Server struct {
	grpcServer *grpc.Server
	grpcLis    net.Listener
	httpServer *http.Server
}

// New creates and configures a new Server instance
func New(cfg *config.Config) (*Server, error) {

	// dependencies
	uc := service.NewInspectionUC(
		builddataclient.NewNHTSA(),
		msrpdataclient.NewMock(),
	)

	// gRPC handler
	grpcSrv := grpc.NewServer()
	pb.RegisterInspectionServiceServer(grpcSrv, grpcDelivery.NewInspectionServer(uc))
	reflection.Register(grpcSrv)
	lis, err := net.Listen("tcp", cfg.GrpcAddress)
	if err != nil {
		return nil, err
	}

	// HTTP gateway mux
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := pb.RegisterInspectionServiceHandlerFromEndpoint(
		context.Background(), mux, cfg.GrpcAddress, opts,
	); err != nil {
		return nil, err
	}
	httpSrv := &http.Server{
		Addr:              cfg.HttpAddress,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// create server
	return &Server{
		grpcServer: grpcSrv,
		grpcLis:    lis,
		httpServer: httpSrv,
	}, nil
}

// Start runs both HTTP and gRPC servers
func (s *Server) Start(ctx context.Context) error {

	// gRPC
	go func() {
		logger.Log.Info("starting gRPC server", slog.String("addr", s.grpcLis.Addr().String()))
		if err := s.grpcServer.Serve(s.grpcLis); err != nil {
			logger.Log.Error("grpc server error", slog.String("err", err.Error()))
		}
	}()

	// HTTP
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		if err := pb.RegisterInspectionServiceHandlerFromEndpoint(ctx, mux, s.grpcLis.Addr().String(), opts); err != nil {
			logger.Log.Error("gateway error", slog.String("err", err.Error()))
		}
		s.httpServer.Handler = mux
		logger.Log.Info("starting HTTP gateway", slog.String("addr", s.httpServer.Addr))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Error("http server error", slog.String("err", err.Error()))
		}
	}()

	return nil
}

// Stop gracefully shuts down both servers
func (s *Server) Stop() error {
	s.grpcServer.GracefulStop()
	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		return err
	}
	return nil
}
