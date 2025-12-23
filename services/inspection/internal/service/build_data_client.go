package service

import "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/entity"

// BuildDataClient defines the interface for fetching build data for vehicles
type BuildDataClient interface {
	Fetch(*entity.Vehicle) error
}
