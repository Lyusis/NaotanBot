package config

type Liver struct {
	RoomId int    `yaml:"room_id"`
	Name   string `yaml:"nickname"`
}

type Livers struct {
	Liver         []Liver `yaml:"liver"`
	GroupId       string  `yaml:"group_id"`
	QQ            string  `yaml:"qq"`
	SaberchanCode string  `yaml:"saberchanCode"`
	CQServer      string  `yaml:"cq_server"`
	Token         string  `yaml:"token"`
	WorkerCount   int     `yaml:"worker_count"`
	Wait          int     `yaml:"waiting"`
}
