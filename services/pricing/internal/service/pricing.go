package service

import (
	"context"

	"github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/entity"
)

// Pricing usecase defines the interface for pricing-related business logic
type PricingUsecase interface {
	GetRecommendedPrice(ctx context.Context, v *entity.Vehicle) error
}

// pricingUsecase is the implementation of PricingUsecase interface
type pricingUsecase struct {
	inspector InspectionClient
}

// NewPricingUC is the constructor for pricingUsecase
func NewPricingUC(client InspectionClient) *pricingUsecase {
	return &pricingUsecase{inspector: client}
}

// GetRecommendedPrice calculates the recommended price for a vehicle
func (uc *pricingUsecase) GetRecommendedPrice(ctx context.Context, v *entity.Vehicle) error {

	// Validate the vehicle data
	if err := v.Validate(); err != nil {
		return err
	}

	// Fetch MRSP
	msrp, err := uc.inspector.GetMsrp(ctx, v.VIN)
	if err != nil {
		return err
	}
	v.Msrp = msrp

	// Calculate the price
	v.CalcPrice()
	return nil
}
