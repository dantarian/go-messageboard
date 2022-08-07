package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type serverConfig struct {
	Port int
}

type Config struct {
	Server *serverConfig
}

func NewConfig() (*Config, error) {
	defaults()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("fatal error reading config file: %w", err)
		}
	}

	if err := validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to process config: %w", err)
	}

	return &config, nil
}

func defaults() {
	viper.SetDefault("server.port", "8080")
}

func validate() error {
	var errors []string

	port := viper.GetInt("server.port")
	if port < 1 || port > 65535 {
		errors = append(errors, "server.port must be in the range 1-65535")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "; "))
	}
	return nil
}
