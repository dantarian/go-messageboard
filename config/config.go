package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type serverConfig struct {
	Port int
}

type dbConfig struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Config struct {
	Server   *serverConfig
	Database *dbConfig
}

func NewConfig() (*Config, error) {
	defaults()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("boards")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.BindEnv("database.password")
	viper.AutomaticEnv()

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
	viper.SetDefault("database.type", "memory")
	viper.SetDefault("database.port", 5432)
}

func validate() error {
	var errors []string

	serverPort := viper.GetInt("server.port")
	if serverPort < 1 || serverPort > 65535 {
		errors = append(errors, "server.port must be in the range 1-65535")
	}

	databaseType := viper.GetString("database.type")
	if databaseType != "memory" && databaseType != "postgres" {
		errors = append(errors, "database.type must be one of 'memory' or 'postgres'")
	}

	databasePort := viper.GetInt("database.port")
	if databaseType == "postgres" && (databasePort < 1 || databasePort > 65535) {
		errors = append(errors, "database.port must be in range 1-65535")
	}

	databasePassword := viper.GetString("database.password")
	if databaseType == "postgres" && databasePassword == "" {
		errors = append(errors, "database.password must be supplied when database.type is 'postgres'")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "; "))
	}
	return nil
}
