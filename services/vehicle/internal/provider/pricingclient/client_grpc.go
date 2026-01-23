package pricingclient

import "github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"

// PricingClientMock is a mock implementation of the PricingProvider interface for testing purposes
type PricingClientMock struct {
	Data *entity.Vehicle
	Err  error
}

// NewMock creates a new PricingClientMock instance
func NewMock(err error) *PricingClientMock {
	return &PricingClientMock{
		Err: err,
	}
}

// GetRecommendedPrice simulates fetching recommended price for a vehicle
func (m *PricingClientMock) GetRecommendedPrice(v *entity.Vehicle) (uint64, error) {
	return 99_000, m.Err
}
