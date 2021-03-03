package config

// Config : Default parameter that taken from environment variables
type Config struct {
	BaseURL  string
	Username string
	Password string
}

// NewConfig : New returns a new Config struct
func NewConfig() *Config {
	LoadEnv()
	return &Config{
		BaseURL:  GetEnv("BaseURL", ""),
		Username: GetEnv("UserName", ""),
		Password: GetEnv("Password", ""),
	}
}
