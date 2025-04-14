package app

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		App
		Database
		Server
	}

	App struct {
		Name        string
		Version     string
		Environment string
	}

	Database struct {
		Host     string
		Port     string
		Name     string
		Username string
		Password string
	}

	Server struct {
		BaseURL string
		Port    string
	}
)

func LoadAppConfig() *Config {
	viper.AutomaticEnv()

	viper.AddConfigPath("./config")
	viper.SetConfigName("debug")
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
