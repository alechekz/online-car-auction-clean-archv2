package service

import (
	"context"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/repository"
)

// VehicleUsecase defines the interface for vehicle-related business logic
type VehicleUsecase interface {
	Create(ctx context.Context, cv *entity.Vehicle) error
	Get(ctx context.Context, vin string) (*entity.Vehicle, error)
	Update(ctx context.Context, v *entity.Vehicle) error
	Delete(ctx context.Context, vin string) error
	List(ctx context.Context) ([]*entity.Vehicle, error)
	Fetch(v *entity.Vehicle) error
}

// vehicleUsecase is the implementation of VehicleUsecase interface
type vehicleUsecase struct {
	repo       repository.VehicleRepository
	inspection InspectionClient
	pricing    PricingClient
}

// NewVehicleUC is the constructor for vehicleUsecase
func NewVehicleUC(r repository.VehicleRepository, i InspectionClient, p PricingClient) *vehicleUsecase {
	return &vehicleUsecase{
		repo:       r,
		inspection: i,
		pricing:    p,
	}
}

// Fetch fetches all necessary data for vehicle processing
func (uc *vehicleUsecase) Fetch(v *entity.Vehicle) error {

	// Fetch build data and merge with user's vehicle data
	bd, err := uc.inspection.GetBuildData(v.VIN)
	if err != nil {
		return err
	}
	if v.Brand == "" {
		v.Brand = bd.Brand
	}
	if v.Engine == "" {
		v.Engine = bd.Engine
	}
	if v.Transmission == "" {
		v.Transmission = bd.Transmission
	}
	v.MSRP = bd.MSRP

	// Perform vehicle inspection to get the grade
	if err := uc.inspection.InspectVehicle(v); err != nil {
		return err
	}

	// Calculate the price based on MSRP and grade
	price, err := uc.pricing.GetRecommendedPrice(v)
	if err != nil {
		return err
	}
	v.Price = price
	return nil
}

// Create creates a new vehicle record
func (uc *vehicleUsecase) Create(ctx context.Context, v *entity.Vehicle) error {

	// Validate the vehicle data
	if err := v.Validate(); err != nil {
		return err
	}

	// Fetch all necessary data for vehicle processing
	if err := uc.Fetch(v); err != nil {
		return err
	}

	// Save the vehicle record
	return uc.repo.Save(ctx, v)
}

// Get retrieves a vehicle by its VIN
func (uc *vehicleUsecase) Get(ctx context.Context, vin string) (*entity.Vehicle, error) {
	return uc.repo.FindByVIN(ctx, vin)
}

// Update updates an existing vehicle record
func (uc *vehicleUsecase) Update(ctx context.Context, v *entity.Vehicle) error {

	// Validate the vehicle data
	if err := v.Validate(); err != nil {
		return err
	}

	// Fetch all necessary data for vehicle processing
	if err := uc.Fetch(v); err != nil {
		return err
	}

	// Update the vehicle record
	if err := uc.repo.Update(ctx, v); err != nil {
		return err
	}
	return nil
}

// Delete deletes a vehicle by its VIN
func (uc *vehicleUsecase) Delete(ctx context.Context, vin string) error {
	return uc.repo.Delete(ctx, vin)
}

// List lists all vehicles
func (uc *vehicleUsecase) List(ctx context.Context) ([]*entity.Vehicle, error) {
	return uc.repo.List(ctx)
}
