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
	WorkerCount   int
	Waiting       int
}
