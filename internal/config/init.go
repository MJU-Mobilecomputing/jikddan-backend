package config

import (
	"errors"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	config *Config
	once   sync.Once
)

func InitConfig() (*Config, error) {
	if config == nil {
		once.Do(func() {
			err := setConfig()
			if err != nil {
				log.Fatal(err)
			}
		})
	} else {
		return nil, errors.New("config is already set")
	}
	return config, nil
}

func setConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}
