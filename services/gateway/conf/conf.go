package conf

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/uim/logger"
)

type Config struct {
	PublicPort    int      `envconfig:"publicPort"`
	ServiceID     string   `envconfig:"serviceId"`
	ServiceName   string   `envconfig:"serviceName"`
	Namespace     string   `envconfig:"namespace"`
	Listen        string   `envconfig:"listen"`
	PublicAddress string   `envconfig:"publicAddress"`
	ConsulURL     string   `envconfig:"consulURL"`
	Tags          []string `envconfig:"tags"`
}

// Init config.
func Init(file string) (*Config, error) {
	viper.SetConfigFile(file)
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/conf")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config file not found: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	logger.Info(config)

	return &config, nil
}
