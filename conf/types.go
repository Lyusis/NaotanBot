package conf

type Liver struct {
	Nickname string
	RoomId   int
}

type Addr struct {
	IP   string
	Port int
}

type Configuration struct {
	Livers        []Liver
	CQReceiver    Addr
	CQSendDest    Addr
	SaberchanCode string
	GroupId       string
	QQ            string
	Token         string
	Announcement  string
	Quit          string
	WorkerCount   int
	Waiting       int
}

const (
	LiversToml        = "livers"
	NicknameToml      = "nickname"
	RoomIdToml        = "roomid"
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
)
