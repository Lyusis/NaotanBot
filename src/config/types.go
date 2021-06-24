package config

type Liver struct {
	RoomId int    `yaml:"room_id"`
	Name   string `yaml:"nickname"`
}

type Livers struct {
	Items         []Liver `yaml:"liver"`
	GroupId       string  `yaml:"group_id"`
	SaberchanCode string  `yaml:"saberchanCode"`
	CQServer      string  `yaml:"cq_server"`
	WorkerCount   int     `yaml:"worker_count"`
	Wait          int     `yaml:"waiting"`
}
