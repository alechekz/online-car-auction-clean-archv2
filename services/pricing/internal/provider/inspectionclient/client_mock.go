package inspectionclient

import (
	"context"

	"github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/entity"
)

// InspectionClientMock is a mock implementation of the InspectionClient interface for testing purposes
type InspectionClientMock struct {
	Data *entity.Vehicle
	Err  error
}

// NewMock creates a new InspectionClientMock instance
func NewMock(msrp uint64, err error) *InspectionClientMock {
	return &InspectionClientMock{
		Data: &entity.Vehicle{
			Msrp: msrp,
		},
		Err: err,
	}
}

// GetMsrp simulates fetching build data for a vehicle
func (m *InspectionClientMock) GetMsrp(ctx context.Context, vin string) (uint64, error) {
	if m.Err != nil {
		return 0, m.Err
	}
	return m.Data.Msrp, nil
}
