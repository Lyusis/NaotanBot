package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"monitor/logger"
)

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

func YAMLParsing(filename string) Livers {
	var items Livers
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error("无法读取yaml配置文件", false, err)
	}

	err = yaml.Unmarshal(yamlFile, &items)
	if err != nil {
		logger.Error("无法解析yaml配置文件", false, err)
	}

	return items
}
