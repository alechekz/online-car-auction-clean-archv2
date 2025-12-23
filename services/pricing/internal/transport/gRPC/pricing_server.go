package grpc

import (
	"context"

	pb "github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/transport/gRPC/proto"

	"github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/service"
)

// PricingServer implements the gRPC server for pricing service
type PricingServer struct {
	pb.UnimplementedPricingServiceServer
	uc service.PricingUsecase
}

// NewPricingServer creates a new PricingServer instance
func NewPricingServer(uc service.PricingUsecase) *PricingServer {
	return &PricingServer{uc: uc}
}

// GetRecommendedPrice retrieves the recommended price for a vehicle by its VIN
func (s *PricingServer) GetRecommendedPrice(ctx context.Context, req *pb.PriceRequest) (*pb.PriceResponse, error) {
	v := &entity.Vehicle{
		VIN:           req.Vin,
		Odometer:      int(req.Odometer),
		Grade:         int(req.Grade),
		ExteriorColor: req.ExteriorColor,
		InteriorColor: req.InteriorColor,
	}
	err := s.uc.GetRecommendedPrice(ctx, v)
	if err != nil {
		return nil, err
	}
	return &pb.PriceResponse{
		Price: v.Price,
	}, nil
}
