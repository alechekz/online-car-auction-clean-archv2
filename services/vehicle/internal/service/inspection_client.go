package service

import "github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"

// InspectionClient defines the interface for fetching vehicle inspection data
type InspectionClient interface {
	InspectVehicle(v *entity.Vehicle) error
	GetBuildData(vin string) (*entity.Vehicle, error)
}
