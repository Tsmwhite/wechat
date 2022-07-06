package config

import (
	"github.com/spf13/viper"
	"wechat/core/log"
)

type Setting struct {
	ViperInstance *viper.Viper
}

func NewSetting(configFileName string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName(configFileName)
	vp.AddConfigPath("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		log.Error.Println("NewSetting ReadInConfig Error:", err)
		return nil, err
	}
	return &Setting{vp}, nil
}

func (set *Setting) UnmarshalKey(key string, configOption interface{}) error {
	if err := set.ViperInstance.UnmarshalKey(key, configOption); err != nil {
		log.Error.Println("Config Setting Error", err)
		return err
	}
	return nil
}
