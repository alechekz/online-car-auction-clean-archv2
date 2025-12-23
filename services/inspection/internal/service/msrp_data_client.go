package service

import "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/entity"

// MSRPDataClient defines the interface for fetching MSRP for vehicles
type MSRPDataClient interface {
	Fetch(v *entity.Vehicle) error
}
