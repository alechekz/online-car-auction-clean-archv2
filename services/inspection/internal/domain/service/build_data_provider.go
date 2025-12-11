package service

import "github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/domain/entity"

// BuildDataProvider defines the interface for fetching build data for vehicles
type BuildDataProvider interface {
	Fetch(*entity.Vehicle) error
}
