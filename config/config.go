package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ChainUrl 		string `mapstructure:"chainUrl"`
	Host 				string `mapstructure:"host"`
	Port 				string `mapstructure:"port"`
	MongoUri		string `mapstructure:"mongoUri"`
}

func NewConfig() *Config {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}