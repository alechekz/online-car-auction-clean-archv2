package testhelpers

import (
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"
)

// NewTestVehiclesBulk is a test valid vehicles bulk instance
func NewTestVehiclesBulk() *model.VehiclesBulk {
	return &model.VehiclesBulk{
		Vehicles: []*entity.Vehicle{NewTestVehicle()},
	}
}
