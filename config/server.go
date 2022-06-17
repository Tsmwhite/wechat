package config

type Message struct {
	FileMax       int
	ContentMaxLen int
	SendTimeout   int
	RevokeTimeout int
}

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

var (
	ServerEnv  = &Server{}
	WebSrvEnv  = &Server{}
	RedisEnv   = &Redis{}
	DBEnv      = &Database{}
	MessageEnv = &Message{}
	MailEnv    = &Mail{}
)

func SetupServer() error {
	setting, err := NewSetting("config")
	if err != nil {
		return err
	}
	err = setting.UnmarshalKey("Server", ServerEnv)
	if err != nil {
		return err
	}
	err = setting.UnmarshalKey("WebServer", WebSrvEnv)
	if err != nil {
		return err
	}
	err = setting.UnmarshalKey("Redis", RedisEnv)
	if err != nil {
		return err
	}
	err = setting.UnmarshalKey("Message", MessageEnv)
	if err != nil {
		return err
	}
	err = setting.UnmarshalKey("Database", DBEnv)
	if err != nil {
		return err
	}
	err = setting.UnmarshalKey("Mail", MailEnv)
	if err != nil {
		return err
	}
	//fmt.Println("server", ServerEnv)
	//fmt.Println("message", MessageEnv)
	//fmt.Println("Database", DBEnv)
	//fmt.Println("Mail", MailEnv)
	return nil
}
