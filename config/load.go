package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Load pulls the config data from the config file
func Load() (*Configuration, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reding config file: %w", err)
	}

	config := Configuration{}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarshalling config file: %w", err)
	}

	validator := validator.New()
	if err := validator.Struct(config); err != nil {
		return nil, fmt.Errorf("invalid config file: %w", err)
	}

	return &config, nil
}
