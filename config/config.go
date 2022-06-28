package config

import (
	"github.com/spf13/viper"
)

func NewSetting(configFileName string) (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName(configFileName)
	vp.AddConfigPath("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return vp, nil
}
