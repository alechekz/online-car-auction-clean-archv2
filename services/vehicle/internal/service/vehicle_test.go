package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	mocks "github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/mock"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/provider/inspectionclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/provider/pricingclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/service"
)

// test is a struct for vehicle service tests
type test struct {
	name    string
	data    func() *entity.Vehicle
	isValid bool
}

// newTestVehicle is a test valid vehicle instance
func newTestVehicle() *entity.Vehicle {
	return &entity.Vehicle{
		VIN:          "1HGCM82633A123456",
		Year:         2020,
		Odometer:     15000,
		MSRP:         25000.00,
		Brand:        "Kia",
		Engine:       "1.8L",
		Transmission: "Automatic",
	}
}

// newTestUC is a helper function to create a Vehicleservice instance for testing
func newTestUC() service.VehicleUsecase {
	v := newTestVehicle()
	repo := mocks.NewVehiclesRepository()
	repo.On("Save", mock.AnythingOfType("*entity.Vehicle")).Return(nil)
	repo.On("FindByVIN", mock.AnythingOfType("string")).Return(v, nil)
	repo.On("Update", mock.AnythingOfType("*entity.Vehicle")).Return(nil)
	repo.On("Delete", mock.AnythingOfType("string")).Return(nil)
	repo.On("List").Return([]*entity.Vehicle{v, v}, nil)

	inspection := inspectionclient.NewMock(v, nil)
	pricing := pricingclient.NewMock(nil)
	return service.NewVehicleUC(repo, inspection, pricing)
}

// TestVehicleservice_CreateVehicle tests the CreateVehicle method of the Vehicleservice struct
func TestVehicleservice_CreateVehicle(t *testing.T) {

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
			err := uc.Create(test.data())
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
	uc := newTestUC()
	v := newTestVehicle()

	// Valid case
	t.Run("existing vehicle", func(t *testing.T) {
		got, err := uc.Get(v.VIN)
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
				return newTestVehicle()
			},
			isValid: true,
		},
		{
			name: "invalid vehicle, too old year",
			data: func() *entity.Vehicle {
				v := newTestVehicle()
				v.Year = 1800
				return v
			},
			isValid: false,
		},
	}

	// Prepare
	uc := newTestUC()
	v := newTestVehicle()
	err := uc.Create(v)
	assert.NoError(t, err)

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := uc.Update(test.data())
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
	uc := newTestUC()
	v := newTestVehicle()

	// Valid case
	t.Run("delete existing vehicle", func(t *testing.T) {
		err := uc.Delete(v.VIN)
		assert.NoError(t, err)
	})
}

// TestVehicleservice_ListVehicles tests the ListVehicles method of the Vehicleservice struct
func TestVehicleservice_ListVehicles(t *testing.T) {

	// Prepare
	uc := newTestUC()

	// List with vehicles case
	t.Run("list with multiple vehicles", func(t *testing.T) {
		vehicles, err := uc.List()
		assert.NoError(t, err)
		assert.Len(t, vehicles, 2)
	})
}
