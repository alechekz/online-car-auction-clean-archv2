package config

import (
	"os"
)

// Config holds the configuration for the Vehicle Service server
type Config struct {
	Address       string
	InspectionURL string
	PricingURL    string
	DatabaseURL   string
}

// New creates a new server configuration with default values
func New() *Config {
	cfg := &Config{
		Address:       ":6061",
		InspectionURL: ":6063",
		PricingURL:    ":6065",
	}
	if os.Getenv("VEHICLE_URL") != "" {
		cfg.Address = os.Getenv("VEHICLE_URL")
	}
	if os.Getenv("INSPECTION_URL") != "" {
		cfg.InspectionURL = os.Getenv("INSPECTION_URL")
	}
	if os.Getenv("PRICING_URL") != "" {
		cfg.PricingURL = os.Getenv("PRICING_URL")
	}
	if os.Getenv("VEHICLE_DB") != "" {
		cfg.DatabaseURL = os.Getenv("VEHICLE_DB")
	}
	return cfg
}
