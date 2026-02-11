package repository

import (
	"context"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"
)

// VehicleRepository defines the interface for vehicle data operations
type VehicleRepository interface {
	Save(ctx context.Context, v *entity.Vehicle) error
	FindByVIN(ctx context.Context, vin string) (*entity.Vehicle, error)
	Update(ctx context.Context, v *entity.Vehicle) error
	Delete(ctx context.Context, vin string) error
	List(ctx context.Context) ([]*entity.Vehicle, error)
	SaveBulk(vb *model.VehiclesBulk) error
	UpdateBulk(vb *model.VehiclesBulk) error
}
