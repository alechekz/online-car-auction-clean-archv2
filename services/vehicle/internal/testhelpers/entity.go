package testhelpers

import "github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"

// NewTestVehicle is a test valid vehicle instance
func NewTestVehicle() *entity.Vehicle {
	return &entity.Vehicle{
		VIN:      "1HGBH41JXMN109186",
		Year:     2022,
		Odometer: 12000,
	}
}
