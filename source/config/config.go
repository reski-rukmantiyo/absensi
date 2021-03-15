package config

// Config : Default parameter that taken from environment variables
type Config struct {
	BaseURL     string
	Username    string
	Password    string
	CacheServer string
	CacheDB     int
	Delay       int
	Port        int
}

// NewConfig : New returns a new Config struct
func NewConfig() *Config {
	LoadEnv()
	return &Config{
		BaseURL:     GetEnv("base_url", ""),
		Username:    GetEnv("username", ""),
		Password:    GetEnv("password", ""),
		CacheServer: GetEnv("cache_server", ""),
		CacheDB:     GetEnvAsInt("cache_db", 0),
		Delay:       GetEnvAsInt("delay", 30),
		Port:        GetEnvAsInt("port", 2112),
	}
}
