package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/provider/inspectionclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/service"
)

// test is a struct for vehicle usecase tests
type test struct {
	name    string
	data    func() *entity.Vehicle
	isValid bool
}

// newTestVehicle is a test valid vehicle instance
func newTestVehicle() *entity.Vehicle {
	return &entity.Vehicle{
		VIN:      "1HGCM82633A123456",
		Odometer: 15000,
		Grade:    47,
	}
}

// newTestUC is a helper function to create a PricingUsecase instance for testing
func newTestUC() service.PricingUsecase {
	inspector := inspectionclient.NewMock(99_000, nil)
	return service.NewPricingUC(inspector)
}

// TestPricingUsecase_GetRecommendedPrice tests the GetRecommendedPrice method of the PricingUsecase struct
func TestPricingUsecase_GetRecommendedPrice(t *testing.T) {

	// Define test cases
	tests := []test{
		{
			name: "valid vehicle",
			data: func() *entity.Vehicle {
				return newTestVehicle()
			},
			isValid: true,
		},
		{
			name: "invalid VIN",
			data: func() *entity.Vehicle {
				v := newTestVehicle()
				v.VIN = "123"
				return v
			},
			isValid: false,
		},
	}

	uc := newTestUC()

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := uc.GetRecommendedPrice(context.Background(), test.data())
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
