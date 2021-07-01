package conf

type Liver struct {
	Nickname string
	RoomId   int
}

type Configuration struct {
	Livers        []Liver
	GroupId       string
	QQ            string
	SaberchanCode string
	CQServer      string
	Token         string
	Announcement  string
	WorkerCount   int
	Waiting       int
}
