package config

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
	Router      string
	Port        string
	PingTimeOut int
}

type WebServer struct {
	Port string
	Tls  bool
}

type Message struct {
	FileMax       int
	ContentMaxLen int
	SendTimeout   int
	RevokeTimeout int
}

var (
	ServerEnv  = &Server{}
	WebSrvEnv  = &WebServer{}
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
	err := LoadMod()
	if err != nil {
		return err
	}
	configFile := ""
	if IsPro() {
		configFile = "config.pro"
	} else {
		configFile = "config.dev"
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
