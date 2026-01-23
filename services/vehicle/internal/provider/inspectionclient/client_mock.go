package inspectionclient

import (
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
)

// InspectionClientMock is a mock implementation of the InspectionClient interface for testing purposes
type InspectionClientMock struct {
	Data *entity.Vehicle
	Err  error
}

// NewMock creates a new InspectionClientMock instance
func NewMock(v *entity.Vehicle, err error) *InspectionClientMock {
	return &InspectionClientMock{
		Data: v,
		Err:  err,
	}
}

// GetBuildData simulates fetching build data for a vehicle
func (m *InspectionClientMock) GetBuildData(vin string) (*entity.Vehicle, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Data, nil
}

// InspectVehicle simulates vehicle inspection
func (m *InspectionClientMock) InspectVehicle(v *entity.Vehicle) error {
	return m.Err
}
