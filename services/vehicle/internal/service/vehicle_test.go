package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/testhelpers"
)

// test is a struct for vehicle service tests
type test struct {
	name    string
	data    func() *entity.Vehicle
	isValid bool
}

// TestVehicleservice_CreateVehicle tests the CreateVehicle method of the Vehicleservice struct
func TestVehicleservice_CreateVehicle(t *testing.T) {

	// Define test cases
	tests := []test{
		{
			name: "valid vehicle",
			data: func() *entity.Vehicle {
				return testhelpers.NewTestVehicle()
			},
			isValid: true,
		},
		{
			name: "invalid VIN",
			data: func() *entity.Vehicle {
				v := testhelpers.NewTestVehicle()
				v.VIN = "123"
				return v
			},
			isValid: false,
		},
	}

	uc := testhelpers.NewTestVehicleUC()

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, stop := context.WithTimeout(context.Background(), 10*time.Second)
			defer stop()
			err := uc.Create(ctx, test.data())
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

// TestVehicleservice_GetVehicle tests the GetVehicle method of the Vehicleservice struct
func TestVehicleservice_GetVehicle(t *testing.T) {

	// Prepare
	uc := testhelpers.NewTestVehicleUC()
	v := testhelpers.NewTestVehicle()

	// Valid case
	t.Run("existing vehicle", func(t *testing.T) {
		ctx, stop := context.WithTimeout(context.Background(), 10*time.Second)
		defer stop()
		got, err := uc.Get(ctx, v.VIN)
		assert.NoError(t, err)
		assert.Equal(t, v.VIN, got.VIN)
	})
}

// TestVehicleservice_UpdateVehicle tests the UpdateVehicle method of the Vehicleservice struct
func TestVehicleservice_UpdateVehicle(t *testing.T) {

	// Define test cases
	tests := []test{
		{
			name: "valid vehicle",
			data: func() *entity.Vehicle {
				return testhelpers.NewTestVehicle()
			},
			isValid: true,
		},
		{
			name: "invalid vehicle, too old year",
			data: func() *entity.Vehicle {
				v := testhelpers.NewTestVehicle()
				v.Year = 1800
				return v
			},
			isValid: false,
		},
	}

	// Prepare
	uc := testhelpers.NewTestVehicleUC()
	v := testhelpers.NewTestVehicle()
	ctx, stop := context.WithTimeout(context.Background(), 10*time.Second)
	defer stop()
	err := uc.Create(ctx, v)
	assert.NoError(t, err)

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := uc.Update(ctx, test.data())
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

// TestVehicleservice_DeleteVehicle tests the DeleteVehicle method of the Vehicleservice struct
func TestVehicleservice_DeleteVehicle(t *testing.T) {

	// Prepare
	uc := testhelpers.NewTestVehicleUC()
	v := testhelpers.NewTestVehicle()

	// Valid case
	t.Run("delete existing vehicle", func(t *testing.T) {
		ctx, stop := context.WithTimeout(context.Background(), 10*time.Second)
		defer stop()
		err := uc.Delete(ctx, v.VIN)
		assert.NoError(t, err)
	})
}

// TestVehicleservice_ListVehicles tests the ListVehicles method of the Vehicleservice struct
func TestVehicleservice_ListVehicles(t *testing.T) {

	// Prepare
	uc := testhelpers.NewTestVehicleUC()

	// List with vehicles case
	t.Run("list with multiple vehicles", func(t *testing.T) {
		ctx, stop := context.WithTimeout(context.Background(), 10*time.Second)
		defer stop()
		vehicles, err := uc.List(ctx)
		assert.NoError(t, err)
		assert.Len(t, vehicles, 2)
	})
}
