package pricingclient

import (
	"context"
	"time"

	pb "github.com/alechekz/online-car-auction-clean-archv2/gen/pricing/v1"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// PricingGRPCClient is a gRPC client for the Pricing Service
type PricingGRPCClient struct {
	client pb.PricingServiceClient
	conn   *grpc.ClientConn
}

// New creates a new PricingGRPCClient instance
func New(address string) (*PricingGRPCClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewPricingServiceClient(conn)
	return &PricingGRPCClient{client: client, conn: conn}, nil
}

// Close closes the gRPC connection
func (c *PricingGRPCClient) Close() error {
	return c.conn.Close()
}

// GetRecommendedPrice sends price request to the Pricing Service
func (c *PricingGRPCClient) GetRecommendedPrice(v *entity.Vehicle) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req := &pb.PriceRequest{
		Vin:           v.VIN,
		Odometer:      v.Odometer,
		Grade:         int32(v.Grade), // nolint:gosec
		ExteriorColor: v.ExteriorColor,
		InteriorColor: v.InteriorColor,
	}
	resp, err := c.client.GetRecommendedPrice(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Price, nil
}
