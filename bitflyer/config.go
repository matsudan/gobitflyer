package bitflyer

import "os"

// A Config provides api configuration for api clients.
type Config struct {
	Region      string
	Credentials Credentials
	// HTTPClient  HTTPClient
	// EndpointResolver EndpointResolver
	// Retryer Retryer
}

// NewConfig returns a new Config pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
func NewConfig() *Config {
	return &Config{}
}

// LoadConfig returns a Config pointer with API credentials fetch from local environmental variables.
func LoadConfig() *Config {
	apiKey := os.Getenv("BITFLYER_API_KEY")
	apiSecret := os.Getenv("BITFLYER_API_SECRET")

	return &Config{
		Credentials: Credentials{
			APIKey: apiKey,
			APISecret: apiSecret,
		},
	}
}