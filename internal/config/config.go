package config

type Config struct {
}

func Phase() (*Config, error) {
	return &Config{}, nil
}
