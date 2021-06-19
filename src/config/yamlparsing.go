package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"monitor/logger"
)

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
