package service

import "github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"

// PricingClient defines the interface for fetching vehicle pricing data
type PricingClient interface {
	GetRecommendedPrice(v *entity.Vehicle) (uint64, error)
}
