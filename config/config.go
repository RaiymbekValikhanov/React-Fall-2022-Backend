package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName    string
	HttpPort   string
	LogLevel   string
	RulesPath  string
	ExamsPath  string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig(path string) *Config {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err)
	}

	return cfg
}
