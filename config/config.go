package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	DBHost string `mapstructure:"dbhost"`
	DBPort string `mapstructure:"dbport"`
	DBUser string `mapstructure:"dbuser"`
	DBPassword string `mapstructure:"dbpassword"`
	DBName string `mapstructure:"dbname"`
}

func NewConfig() Config {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}