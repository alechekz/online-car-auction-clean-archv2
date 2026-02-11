package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"
)

// MockVehiclesRepository is a mock implementation of VehiclesRepository for testing
type MockVehiclesRepository struct {
	mock.Mock
}

// NewVehiclesRepo creates a new instance of MockVehiclesRepository
func NewVehiclesRepo() *MockVehiclesRepository {
	return &MockVehiclesRepository{}
}

// Save saves a vehicle record
func (m *MockVehiclesRepository) Save(ctx context.Context, v *entity.Vehicle) error {
	args := m.Called(v)
	return args.Error(0)
}

// Update updates a vehicle record
func (m *MockVehiclesRepository) Update(ctx context.Context, v *entity.Vehicle) error {
	args := m.Called(v)
	return args.Error(0)
}

// FindByVIN finds a vehicle by its VIN
func (m *MockVehiclesRepository) FindByVIN(ctx context.Context, vin string) (*entity.Vehicle, error) {
	args := m.Called(vin)
	return args.Get(0).(*entity.Vehicle), args.Error(1)
}

// Delete deletes a vehicle by its VIN
func (m *MockVehiclesRepository) Delete(ctx context.Context, vin string) error {
	args := m.Called(vin)
	return args.Error(0)
}

// List lists all vehicles
func (m *MockVehiclesRepository) List(ctx context.Context) ([]*entity.Vehicle, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Vehicle), args.Error(1)
}

// SaveBulk saves multiple vehicle records in bulk
func (m *MockVehiclesRepository) SaveBulk(vb *model.VehiclesBulk) error {
	args := m.Called(vb)
	return args.Error(0)
}

// UpdateBulk updates multiple vehicle records in bulk
func (m *MockVehiclesRepository) UpdateBulk(vb *model.VehiclesBulk) error {
	args := m.Called(vb)
	return args.Error(0)
}
