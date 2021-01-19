package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database Dbconfig
}

type Dbconfig struct {
	Host     string
	Port     string
	Db       string
	Username string
	Password string
	SSLMode  string
}

func ReadConfig() *Config {
	viper.SetConfigFile("shurly-config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarshaling config: %s \n", err))
	}

	return &config
}
