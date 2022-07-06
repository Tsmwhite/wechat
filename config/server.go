package config

import (
	"wechat/env"
)

type Database struct {
	Host,
	Port,
	UserName,
	Password,
	DatabaseName string
	MaxIdleConns,
	MaxOpenConns,
	ConnMaxLifetime int
	Dsn string
}

type Mail struct {
	Host,
	Port,
	Account,
	Password string
}

type Redis struct {
	Addr     string
	Port     int
	Password string
}

type Server struct {
	Router,
	Port string
	PingTimeOut int
}

type WebServer struct {
	Port string
}

type Message struct {
	FileMax       int
	ContentMaxLen int
	SendTimeout   int
	RevokeTimeout int
}

var (
	ServerEnv  = &Server{}
	WebSrvEnv  = &Server{}
	RedisEnv   = &Redis{}
	DBEnv      = &Database{}
	MessageEnv = &Message{}
	MailEnv    = &Mail{}
)

var configMap = map[string]interface{}{
	"Server":    ServerEnv,
	"WebServer": WebSrvEnv,
	"Redis":     RedisEnv,
	"Database":  DBEnv,
	"Mail":      MailEnv,
	"Message":   MessageEnv,
}

func SetupServer() error {
	configFile := ""
	if env.Debug {
		configFile = "config.dev"
	} else {
		configFile = "config"
	}
	setting, err := NewSetting(configFile)
	if err != nil {
		return err
	}
	for key, option := range configMap {
		if err = setting.UnmarshalKey(key, option); err != nil {
			return err
		}
	}
	return nil
}
