package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
)

// VehiclesBulk represents a bulk operation on vehicles
type VehiclesBulk struct {
	Vehicles []*entity.Vehicle `json:"vehicles"`
}

// Validate checks if the VehiclesBulk data is valid
func (vb *VehiclesBulk) Validate() error {
	return validation.ValidateStruct(
		vb,
		validation.Field(
			&vb.Vehicles,
			validation.Required,
			validation.Each(
				validation.Required,
			),
		),
	)
}
