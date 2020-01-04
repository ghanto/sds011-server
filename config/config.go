package config

// Config structure of config
type Config struct {
	DatabaseDSN string `env:"DATABASE_DSN"`
}

// ApplicationConfig configuration of the app
var ApplicationConfig = Config{}
