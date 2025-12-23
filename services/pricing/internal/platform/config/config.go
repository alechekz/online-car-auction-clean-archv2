package config

import "os"

// Config holds the configuration for the Inspection Service server
type Config struct {
	HttpAddress   string
	GrpcAddress   string
	DatabaseURL   string
	InspectionURL string
}

// New creates a new server configuration with default values
func New() *Config {
	cfg := &Config{
		HttpAddress:   ":6064",
		GrpcAddress:   ":6065",
		InspectionURL: ":6063",
	}
	if os.Getenv("PRICING_HTTP") != "" {
		cfg.HttpAddress = os.Getenv("PRICING_HTTP")
	}
	if os.Getenv("PRICING_GRPC") != "" {
		cfg.GrpcAddress = os.Getenv("PRICING_GRPC")
	}
	if os.Getenv("INSPECTION_URL") != "" {
		cfg.InspectionURL = os.Getenv("INSPECTION_URL")
	}
	return cfg
}
