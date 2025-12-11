package service

import (
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/domain/entity"
)

// InspectionUsecase defines the interface for inspection-related business logic
type InspectionUsecase interface {
	InspectVehicle(v *entity.Vehicle) error
	GetBuildData(vin string) (*entity.Vehicle, error)
}

// inspectionUsecase is the implementation of InspectionUsecase interface
type inspectionUsecase struct {
	dataProvider BuildDataProvider
	msrpProvider MSRPProvider
}

// NewInspectionUC is the constructor for inspectionUsecase
func NewInspectionUC(data BuildDataProvider, msrp MSRPProvider) InspectionUsecase {
	return &inspectionUsecase{dataProvider: data, msrpProvider: msrp}
}

// InspectVehicle inspects a vehicle and creates a new inspection record
func (uc *inspectionUsecase) InspectVehicle(v *entity.Vehicle) error {

	// Validate the inspection data
	if err := v.Validate(); err != nil {
		return err
	}

	// Make inspection (dummy logic for example)
	v.Inspect()
	return nil
}

// GetBuildData retrieves the build data for a vehicle by its VIN
func (uc *inspectionUsecase) GetBuildData(vin string) (*entity.Vehicle, error) {

	// Prepare the vehicle instance and validate VIN
	v := &entity.Vehicle{VIN: vin}
	if err := v.ValidateVIN(); err != nil {
		return nil, err
	}

	// Fetch MSRP data
	if err := uc.msrpProvider.Fetch(v); err != nil {
		return nil, err
	}

	// Fetch build data from the provider
	return v, uc.dataProvider.Fetch(v)
}
