package config

import "go-ecm/internal/goecmagent/options"

type Config struct {
	*options.Options
}

func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{
		opts,
	}, nil
}
