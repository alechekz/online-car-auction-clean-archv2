package testhelpers

import (
	"github.com/stretchr/testify/mock"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/mocks"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/provider/inspectionclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/provider/pricingclient"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/service"
)

// NewTestVehicleUC is a helper function to create a Vehicleservice instance for testing
func NewTestVehicleUC() service.VehicleUsecase {
	v := NewTestVehicle()
	repo := mocks.NewVehiclesRepo()
	repo.On("Save", mock.AnythingOfType("*entity.Vehicle")).Return(nil)
	repo.On("FindByVIN", mock.AnythingOfType("string")).Return(v, nil)
	repo.On("Update", mock.AnythingOfType("*entity.Vehicle")).Return(nil)
	repo.On("Delete", mock.AnythingOfType("string")).Return(nil)
	repo.On("List").Return([]*entity.Vehicle{v, v}, nil)

	inspection := inspectionclient.NewMock(v, nil)
	pricing := pricingclient.NewMock(nil)
	return service.NewVehicleUC(repo, inspection, pricing)
}
