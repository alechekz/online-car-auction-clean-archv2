package grpc

import (
	"context"

	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/domain/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/domain/service"
	pb "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/transport/gRPC/proto"
)

// InspectionServer implements the gRPC server for inspection service
type InspectionServer struct {
	pb.UnimplementedInspectionServiceServer
	uc service.InspectionUsecase
}

// NewInspectionServer creates a new InspectionServer instance
func NewInspectionServer(uc service.InspectionUsecase) *InspectionServer {
	return &InspectionServer{uc: uc}
}

// GetBuildData retrieves the build data for a vehicle by its VIN
func (s *InspectionServer) GetBuildData(ctx context.Context, req *pb.GetBuildDataRequest) (*pb.BuildDataResponse, error) {
	data, err := s.uc.GetBuildData(req.Vin)
	if err != nil {
		return nil, err
	}
	return &pb.BuildDataResponse{
		Vin:          data.VIN,
		Brand:        data.Brand,
		Engine:       data.Engine,
		Transmission: data.Transmission,
		Msrp:         uint64(data.MSRP),
	}, nil
}

// InspectVehicle inspects a vehicle and returns its grade
func (s *InspectionServer) InspectVehicle(ctx context.Context, req *pb.InspectVehicleRequest) (*pb.InspectVehicleResponse, error) {
	v := &entity.Vehicle{
		VIN:             req.Vin,
		Odometer:        int(req.Odometer),
		Year:            int(req.Year),
		StrongScratches: req.StrongScratches,
		SmallScratches:  req.SmallScratches,
		ElectricFail:    req.ElectricFail,
		SuspensionFail:  req.SuspensionFail,
	}
	if err := s.uc.InspectVehicle(v); err != nil {
		return nil, err
	}
	return &pb.InspectVehicleResponse{
		Vin:   req.Vin,
		Grade: int32(v.Grade), //nolint:gosec // always between 1 and 50
	}, nil
}
