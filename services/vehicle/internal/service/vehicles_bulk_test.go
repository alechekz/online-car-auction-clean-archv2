package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/mocks"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/service"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/testhelpers"
)

// newTestVehiclesBulk is a test valid VehiclesBulk instance
func newTestVehiclesBulk() *model.VehiclesBulk {
	return &model.VehiclesBulk{
		Vehicles: []*entity.Vehicle{
			testhelpers.NewTestVehicle(),
			testhelpers.NewTestVehicle(),
			testhelpers.NewTestVehicle(),
			testhelpers.NewTestVehicle(),
			testhelpers.NewTestVehicle(),
			testhelpers.NewTestVehicle(),
			testhelpers.NewTestVehicle(),
		},
	}
}

// TestVehiclesBulkUsecase_Create tests the Create method of VehiclesBulkUsecase
func TestVehiclesBulkUsecase_Create(t *testing.T) {
	mockRepo := mocks.NewVehiclesRepo()
	vehicleUC := testhelpers.NewTestVehicleUC()
	uc := service.NewVehiclesBulkUC(mockRepo, vehicleUC)
	vb := newTestVehiclesBulk()

	// Successful case
	mockRepo.On("SaveBulk", vb).Return(nil)
	t.Run("valid bulk data", func(t *testing.T) {
		assert.NoError(t, uc.Create(vb))
	})

	// Return error on Validate failure
	mockRepo.On("SaveBulk", vb).Return(nil)
	v := vb.Vehicles[0]
	v.VIN = "123"
	t.Run("invalid bulk data", func(t *testing.T) {
		assert.Error(t, uc.Create(vb))
	})

	// Return error on repository failure
	mockRepo.On("SaveBulk", vb).Return(assert.AnError)
	t.Run("repository failure", func(t *testing.T) {
		assert.Error(t, uc.Create(vb))
	})
}

// TestVehiclesBulkUsecase_Update tests the Update method of VehiclesBulkUsecase
func TestVehiclesBulkUsecase_Update(t *testing.T) {
	mockRepo := mocks.NewVehiclesRepo()
	vehicleUC := testhelpers.NewTestVehicleUC()
	uc := service.NewVehiclesBulkUC(mockRepo, vehicleUC)
	vb := newTestVehiclesBulk()

	// Successful case
	mockRepo.On("UpdateBulk", vb).Return(nil)
	t.Run("valid bulk data", func(t *testing.T) {
		assert.NoError(t, uc.Update(vb))
	})

	// Return error on Validate failure
	mockRepo.On("UpdateBulk", vb).Return(nil)
	v := vb.Vehicles[0]
	v.VIN = "123"
	t.Run("invalid bulk data", func(t *testing.T) {
		assert.Error(t, uc.Update(vb))
	})

	// Return error on repository failure
	mockRepo.On("UpdateBulk", vb).Return(assert.AnError)
	t.Run("repository failure", func(t *testing.T) {
		assert.Error(t, uc.Update(vb))
	})
}
