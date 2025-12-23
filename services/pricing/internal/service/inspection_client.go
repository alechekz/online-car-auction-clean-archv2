package service

import "context"

// InspectionClient defines the interface for fetching vehicle build data
type InspectionClient interface {
	GetMsrp(ctx context.Context, vin string) (uint64, error)
}
