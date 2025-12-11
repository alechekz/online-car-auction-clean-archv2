package service

import "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/domain/entity"

// MSRPProvider defines the interface for fetching MSRP for vehicles
type MSRPProvider interface {
	Fetch(v *entity.Vehicle) error
}
