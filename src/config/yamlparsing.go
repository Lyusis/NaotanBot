package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"logger"
)

func YAMLParsing(filename string) Livers {
	var items Livers
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Sugar.Error("无法读取yaml配置文件", logger.FormatTitle("WRONG"), err)
	}

	err = yaml.Unmarshal(yamlFile, &items)
	if err != nil {
		logger.Sugar.Error("无法解析yaml配置文件", logger.FormatTitle("WRONG"), err)
	}

	return items
}
