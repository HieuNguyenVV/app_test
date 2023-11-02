package config

import (
	"app/pkg/db"

	"github.com/spf13/viper"
)

var (
	ConfigFileName string = "config"
	ConfigType     string = "yaml"
	ConfigPath     string = "C:/Users/Lenovo/go/src/app/configs/"
)

type Config struct {
	Postgres db.Postgres `mapstructure:"postgres"`
}

func AutoBindConfig() (*Config, error) {
	vp := viper.New()

	vp.AddConfigPath(ConfigPath)
	vp.SetConfigName(ConfigFileName)
	vp.SetConfigType(ConfigType)

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	conf := Config{}
	if err := vp.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
