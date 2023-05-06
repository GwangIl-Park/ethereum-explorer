package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host 				string `mapstructure:"host"`
	Port 				string `mapstructure:"port"`
	ChainUrl 		string `mapstructure:"chainUrl"`
	MongoUri		string `mapstructure:"mongoUri"`
	StartBlock	int64 `mapstructure:"startBlock"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}