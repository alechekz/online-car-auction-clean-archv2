package model_test

import (
	"testing"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"

	"github.com/stretchr/testify/assert"
)

// vBtest is a struct for vehicles bulk tests
type vBtest struct {
	name    string
	data    func() *model.VehiclesBulk
	isValid bool
}

// newTestVehicle is a test valid vehicle instance
func newTestVehicle() *entity.Vehicle {
	return &entity.Vehicle{
		VIN:      "1HGBH41JXMN109186",
		Year:     2022,
		Odometer: 12000,
	}
}

// newTestVehiclesBulk is a test valid vehicles bulk instance
func newTestVehiclesBulk() *model.VehiclesBulk {
	return &model.VehiclesBulk{
		Vehicles: []*entity.Vehicle{newTestVehicle()},
	}
}

// TestVehiclesBulk_Validate tests the Validate method of the VehiclesBulk struct
func TestVehiclesBulk_Validate(t *testing.T) {
	tests := []vBtest{
		{
			name: "valid vehicles bulk",
			data: func() *model.VehiclesBulk {
				return newTestVehiclesBulk()
			},
			isValid: true,
		},
		{
			name: "empty vehicles slice",
			data: func() *model.VehiclesBulk {
				vb := newTestVehiclesBulk()
				vb.Vehicles = []*entity.Vehicle{}
				return vb
			},
			isValid: false,
		},
		{
			name: "vehicles slice with invalid vehicle",
			data: func() *model.VehiclesBulk {
				vb := newTestVehiclesBulk()
				v := newTestVehicle()
				v.VIN = "123"
				vb.Vehicles = append(vb.Vehicles, v)
				return vb
			},
			isValid: false,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.data().Validate()
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
