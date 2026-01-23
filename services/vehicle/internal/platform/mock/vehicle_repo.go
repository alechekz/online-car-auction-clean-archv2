package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"
)

// MockVehiclesRepository is a mock implementation of VehiclesRepository for testing
type MockVehiclesRepository struct {
	mock.Mock
}

// NewVehiclesRepository creates a new instance of MockVehiclesRepository
func NewVehiclesRepository() *MockVehiclesRepository {
	return &MockVehiclesRepository{}
}

// Save saves a vehicle record
func (m *MockVehiclesRepository) Save(v *entity.Vehicle) error {
	args := m.Called(v)
	return args.Error(0)
}

// Update updates a vehicle record
func (m *MockVehiclesRepository) Update(v *entity.Vehicle) error {
	args := m.Called(v)
	return args.Error(0)
}

// FindByVIN finds a vehicle by its VIN
func (m *MockVehiclesRepository) FindByVIN(vin string) (*entity.Vehicle, error) {
	args := m.Called(vin)
	return args.Get(0).(*entity.Vehicle), args.Error(1)
}

// Delete deletes a vehicle by its VIN
func (m *MockVehiclesRepository) Delete(vin string) error {
	args := m.Called(vin)
	return args.Error(0)
}

// List lists all vehicles
func (m *MockVehiclesRepository) List() ([]*entity.Vehicle, error) {
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
