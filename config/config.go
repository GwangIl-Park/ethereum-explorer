package config

import (
	"ethereum-explorer/logger"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Url        string `mapstructure:"url`
	ChainHttp  string `mapstructure:"chainHttp"`
	ChainWs    string `mapstructure:"chainWs"`
	DbHost     string `mapstructure:"dbHost"`
	DbPort     int32  `mapstructure:"dbPort"`
	DbUser     string `mapstructure:"dbUser"`
	DbPassword string `mapstructure:"dbPassword"`
	DbName     string `mapstructure:"dbName"`
	StartBlock int64  `mapstructure:"startBlock"` //For Test
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	logger.Logger.WithFields(
		logrus.Fields{
			"url":        cfg.Url,
			"chainHttp":  cfg.ChainHttp,
			"chainWs":    cfg.ChainWs,
			"dbHost":     cfg.DbHost,
			"dbPort":     cfg.DbPort,
			"dbUser":     cfg.DbUser,
			"dbPassword": cfg.DbPassword,
			"dbName":     cfg.DbName,
			"startBlock": cfg.StartBlock,
		},
	).Debug("Check Flag")

	return &cfg, nil
}
