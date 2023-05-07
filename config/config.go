package config

import (
	"ethereum-explorer/logger"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Host 				string `mapstructure:"host"`
	Port 				string `mapstructure:"port"`
	ChainHttp 		string `mapstructure:"chainHttp"`
	ChainWs 		string `mapstructure:"chainWs"`
	MongoUri		string `mapstructure:"mongoUri"`
	StartBlock	int64 `mapstructure:"startBlock"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	
	logger.Logger.WithFields(
		logrus.Fields{
				"host":cfg.Host,
				"port":cfg.Port,
				"chainHttp":cfg.ChainHttp,
				"chainWs":cfg.ChainWs,
				"mongoUri":cfg.MongoUri,
				"startBlock":cfg.StartBlock,
			},
	).Debug("Check Flag")

	return &cfg, nil
}