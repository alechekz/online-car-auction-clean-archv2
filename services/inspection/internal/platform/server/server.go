package server

import (
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcDelivery "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/transport/gRPC"
	pb "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/transport/gRPC/proto"

	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/domain/service"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/config"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/platform/logger"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/provider"
)

// Server represents both HTTP and gRPC servers for the Inspection Service
type Server struct {
	grpcServer *grpc.Server
	grpcLis    net.Listener
}

// New creates and configures a new Server instance
func New(cfg *config.Config) (*Server, error) {

	// dependencies
	uc := service.NewInspectionUC(
		provider.NewNHTSABuildDataClient(),
		provider.NewMockMSRPClient(),
	)

	// gRPC handler
	grpcSrv := grpc.NewServer()
	pb.RegisterInspectionServiceServer(grpcSrv, grpcDelivery.NewInspectionServer(uc))
	reflection.Register(grpcSrv)
	lis, err := net.Listen("tcp", cfg.GrpcAddress)
	if err != nil {
		return nil, err
	}

	// create server
	return &Server{
		grpcServer: grpcSrv,
		grpcLis:    lis,
	}, nil
}

// Start runs both HTTP and gRPC servers
func (s *Server) Start() error {
	// gRPC
	go func() {
		logger.Log.Info("starting gRPC server", slog.String("addr", s.grpcLis.Addr().String()))
		if err := s.grpcServer.Serve(s.grpcLis); err != nil {
			logger.Log.Error("grpc server error", slog.String("err", err.Error()))
		}
	}()

	return nil
}

// Stop gracefully shuts down both servers
func (s *Server) Stop() error {
	s.grpcServer.GracefulStop()
	return nil
}
