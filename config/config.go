package config

// Config : Default parameter that taken from environment variables
type Config struct {
	BaseURL string
}

// NewConfig : New returns a new Config struct
func NewConfig() *Config {
	LoadEnv()
	return &Config{
		BaseURL: GetEnv("BASE_URL", ""),
	}
}
