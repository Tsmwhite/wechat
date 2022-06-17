package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

var RootPath string

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

type config struct {
	Message struct {
		Status        map[string]int `yaml:"status"`
		FileMax       int            `yaml:"file_max"`
		ContentMaxLen int            `yaml:"content_max_len"`
		SendTimeout   int            `yaml:"send_time_out"`
		RevokeTimeout int            `yaml:"revoke_time_out"`
	}
	Client struct {
		PingTimeOut int `yaml:"ping_time_out"`
	}
	Server struct {
		Router string `yaml:"router"`
		Port   string `yaml:"port"`
	}
}

const ConfigFile = "/config/config.yaml"

var CONFIG config

func init() {
	//项目根目录
	RootPath, _ = os.Getwd()
	//配置文件路径
	CONFIG_PATH := RootPath + ConfigFile
	//判断配置文件是否存在
	_, err := os.Stat(CONFIG_PATH)
	if err != nil {
		if os.IsNotExist(err) {
			panic("Missing core configuration file：" + CONFIG_PATH)
		}
		panic("Missing core configuration file：" + err.Error())
	}

	//获取配置信息
	data, err := ioutil.ReadFile(CONFIG_PATH)
	if err != nil {
		panic("Error reading configuration file：" + err.Error())
	}

	//转化为Config struct
	err = yaml.Unmarshal(data, &CONFIG)
	if err != nil {
		panic("Please check if the configuration file is written correctly")
	}
}
