package conf

type Addr struct {
	IP   string
	Port int
}
type Redis struct {
	IP       string
	Port     int
	Password string
}

type Configuration struct {
	RedisInfo     Redis
	CQReceiver    Addr
	CQSendDest    Addr
	SaberchanCode string
	Token         string
	Announcement  string
	Quit          string
	WeatherSecret string
	NewsKey       string
	QQ            int
	GroupId       int
	AJun          int
	WorkerCount   int
	Waiting       int
}

const (
	CQReceiverToml    = "cqreceiver"
	CQSendDestToml    = "cqsenddest"
	SaberchanCodeToml = "saberchancode"
	GroupIdToml       = "groupid"
	QQToml            = "qq"
	TokenToml         = "token"
	AnnouncementToml  = "announcement"
	QuitMessageToml   = "quit"
	WorkerCountToml   = "workercount"
	WaitingToml       = "waiting"
	IPToml            = "ip"
	PortToml          = "port"
	RedisInfoToml     = "redisinfo"
	WeatherSecretToml = "weathersecret"
	AJunToml          = "ajun"
	NewsKeyToml       = "newskey"
)
