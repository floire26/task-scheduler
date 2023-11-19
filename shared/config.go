package shared

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"SERVER_PORT"`
	TimeoutDur int    `mapstructure:"TIMEOUT_DUR"`
}

func LoadConfig(path string) *Config {
	cfg := &Config{}

	viper.SetConfigFile(path)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("unable to find the config file: %v", err)
		return nil
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("unable to load the environment: %v", err)
		return nil
	}

	return cfg
}
