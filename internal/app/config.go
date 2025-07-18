package app

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	configName   = "config"
	configType   = "yaml"
	pathToConfig = "./config"
)

const (
	defaultServerHost = "0.0.0.0"
	defaultServerPort = "8080"
)

type appConfig struct {
	Server struct {
		Host string
		Port string
	}
}

func loadConfig() (*appConfig, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(pathToConfig)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("read yaml app config: %w", err)
	}

	var cfg appConfig
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal app config: %w", err)
	}

	cfg = setDefaults(cfg)

	return &cfg, nil
}

func setDefaults(cfg appConfig) appConfig {
	if cfg.Server.Host == "" {
		cfg.Server.Host = defaultServerHost
	}

	if cfg.Server.Port == "" {
		cfg.Server.Port = defaultServerPort
	}

	return cfg
}
