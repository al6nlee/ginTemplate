package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
