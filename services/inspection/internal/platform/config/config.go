package config

import "os"

// Config holds the configuration for the Inspection Service server
type Config struct {
	HttpAddress string
	GrpcAddress string
	DatabaseURL string
}

// New creates a new server configuration with default values
func New() *Config {
	cfg := &Config{
		HttpAddress: ":6062",
		GrpcAddress: ":6063",
	}
	if os.Getenv("INSPECTION_HTTP") != "" {
		cfg.HttpAddress = os.Getenv("INSPECTION_HTTP")
	}
	if os.Getenv("INSPECTION_GRPC") != "" {
		cfg.GrpcAddress = os.Getenv("INSPECTION_GRPC")
	}
	return cfg
}
