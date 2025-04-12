package configs

import (
	"github.com/joaoasantana/e-product-service/pkg/util/configs"
	"github.com/spf13/viper"
)

type Config struct {
	configs.App
	configs.Database
	configs.Server
}

func LoadNewConfig() *Config {
	viper.AddConfigPath("./env")
	viper.AutomaticEnv()

	viper.SetConfigName("config.debug")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
