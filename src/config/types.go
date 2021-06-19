package config

type Liver struct {
	RoomId int    `yaml:"room_id"`
	Name   string `yaml:"nickname"`
}

type Livers struct {
	Items         []Liver `yaml:"liver"`
	Wait          int     `yaml:"waiting"`
	GroupId       string  `yaml:"group_id"`
	SaberchanCode string  `yaml:"saberchanCode"`
	CQServer      string  `yaml:"cq_server"`
}
