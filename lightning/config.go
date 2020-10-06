package lightning

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

// Copy will return a shallow copy of the Config object. If any additional
// configurations are provided they will be merged into the new config returned.
func (c Config) Copy() Config {
	cp := c
	return cp
}
