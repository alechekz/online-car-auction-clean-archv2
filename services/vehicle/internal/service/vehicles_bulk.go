package service

import (
	"golang.org/x/sync/errgroup"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/repository"
)

// VehiclesBulkUsecase defines the interface for bulk vehicle operations
type VehiclesBulkUsecase interface {
	Create(vb *model.VehiclesBulk) error
	Update(vb *model.VehiclesBulk) error
}

// vehiclesBulkUsecase is the implementation of VehiclesBulkUsecase interface
type vehiclesBulkUsecase struct {
	repo      repository.VehicleRepository
	vehicleUC VehicleUsecase
}

// NewVehiclesBulkUC is the constructor for vehiclesBulkUsecase
func NewVehiclesBulkUC(r repository.VehicleRepository, vehicleUC VehicleUsecase) *vehiclesBulkUsecase {
	return &vehiclesBulkUsecase{
		repo:      r,
		vehicleUC: vehicleUC,
	}
}

// fetch fetches all necessary data for vehicle processing
func (uc *vehiclesBulkUsecase) fetch(vb *model.VehiclesBulk) error {
	g := new(errgroup.Group)
	jobs := make(chan *entity.Vehicle, len(vb.Vehicles))
	workers := 5

	for range workers {
		g.Go(func() error {
			for v := range jobs {
				if err := uc.vehicleUC.Fetch(v); err != nil {
					return err
				}
			}
			return nil
		})
	}

	// Enqueue jobs
	go func() {
		for _, v := range vb.Vehicles {
			jobs <- v
		}
		close(jobs)
	}()

	// Wait for all workers to finish and exit
	return g.Wait()
}

// Create creates multiple vehicle records in bulk
func (uc *vehiclesBulkUsecase) Create(vb *model.VehiclesBulk) error {

	// Validate the bulk vehicle data
	if err := vb.Validate(); err != nil {
		return err
	}

	// Fetch all necessary data for each vehicle in bulk concurrently
	if err := uc.fetch(vb); err != nil {
		return err
	}

	// Save the bulk vehicles
	return uc.repo.SaveBulk(vb)
}

// Update updates multiple vehicle records in bulk
func (uc *vehiclesBulkUsecase) Update(vb *model.VehiclesBulk) error {

	// Validate the bulk vehicle data
	if err := vb.Validate(); err != nil {
		return err
	}

	// Fetch all necessary data for each vehicle in bulk concurrently
	if err := uc.fetch(vb); err != nil {
		return err
	}

	// Update the bulk vehicles
	return uc.repo.UpdateBulk(vb)
}
