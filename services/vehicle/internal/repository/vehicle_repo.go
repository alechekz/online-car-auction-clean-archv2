package repository

import (
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"
)

// VehicleRepository defines the interface for vehicle data operations
type VehicleRepository interface {
	Save(v *entity.Vehicle) error
	FindByVIN(vin string) (*entity.Vehicle, error)
	Update(v *entity.Vehicle) error
	Delete(vin string) error
	List() ([]*entity.Vehicle, error)

	SaveBulk(vb *model.VehiclesBulk) error
	UpdateBulk(vb *model.VehiclesBulk) error
}
