package config

// Config : Default parameter that taken from environment variables
type Config struct {
	BaseURL     string
	Username    string
	Password    string
	Picture     string
	Longitude   string
	Lattitude   string
	Description string
	Region      string
	CacheServer string
	CacheDB     int
}

// NewConfig : New returns a new Config struct
func NewConfig() *Config {
	LoadEnv()
	return &Config{
		BaseURL:     GetEnv("BaseURL", ""),
		Username:    GetEnv("UserName", ""),
		Password:    GetEnv("Password", ""),
		Picture:     GetEnv("Picture", ""),
		Longitude:   GetEnv("Longitude", ""),
		Lattitude:   GetEnv("Lattitude", ""),
		Description: GetEnv("Description", ""),
		Region:      GetEnv("Region", ""),
		CacheServer: GetEnv("", ""),
		CacheDB:     GetEnvAsInt("", 0),
	}
}
