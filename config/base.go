package config

import (
	"fmt"
	"github.com/spf13/viper"
	"wechat/core/log"
)

const ModDevelopment = "development"
const ModProduction = "production"

type Setting struct {
	ViperInstance *viper.Viper
}

var mod string
var debug bool
var webTls bool

func IsDev() bool {
	return ModDevelopment == mod
}

func IsPro() bool {
	return ModProduction == mod
}

func Debug() bool {
	return debug
}

func WebTlsOpen() bool {
	return webTls
}

func LoadMod() error {
	setting, err := NewSetting("config")
	if err != nil {
		return err
	}
	debug = setting.ViperInstance.GetBool("Debug")
	mod = setting.ViperInstance.GetString("EnvMod")
	webTls = setting.ViperInstance.GetBool("WebTls")
	fmt.Println("Debug:", debug)
	fmt.Println("EnvMod:", mod)
	fmt.Println("WebTls:", webTls)
	return nil
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
