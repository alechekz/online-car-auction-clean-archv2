package builddataclient

import (
	"github.com/alechekz/online-car-auction-clean-archv2/services/inspection/internal/entity"
)

// MockBuildDataClient is a mock client for fetching build data based on VIN
type MockBuildDataClient struct {
	data map[string]struct {
		Brand        string
		Engine       string
		Transmission string
	}
}

// NewMock creates a new mock client with predefined data
func NewMock() *MockBuildDataClient {
	return &MockBuildDataClient{
		data: map[string]struct {
			Brand        string
			Engine       string
			Transmission string
		}{
			"5YJSA1E26MF168123": {Brand: "Tesla", Engine: "Electric", Transmission: "Automatic"},
			"WBS8M9C59J5G12345": {Brand: "BMW", Engine: "3.0L Turbo", Transmission: "Automatic"},
			"1HGCM82633A004352": {Brand: "Honda", Engine: "2.4L", Transmission: "Manual"},
			"3FA6P0K92HR123456": {Brand: "Ford", Engine: "2.0L EcoBoost", Transmission: "Automatic"},
			"JHMFA16586S123456": {Brand: "Honda", Engine: "1.3L Hybrid", Transmission: "CVT"},
		},
	}
}

// Fetch returns data from the map based on VIN
func (c *MockBuildDataClient) Fetch(v *entity.Vehicle) error {
	if val, ok := c.data[v.VIN]; ok {
		v.Brand = val.Brand
		v.Engine = val.Engine
		v.Transmission = val.Transmission
	} else {
		// default values for unknown VINs
		v.Brand = "MockBrand"
		v.Engine = "MockEngine"
		v.Transmission = "MockTransmission"
	}
	return nil
}
