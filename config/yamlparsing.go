package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/Lyusis/NaotanMonitor/logger"
)

func YAMLParsing(filename string) Livers {
	var items Livers
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Sugar.Error(logger.FormatMsg("Unable to read the YAML configuration file"), logger.FormatTitle("WRONG"), err)
	}

	err = yaml.Unmarshal(yamlFile, &items)
	if err != nil {
		logger.Sugar.Error(logger.FormatMsg("Unable to parse the YAML configuration file"), logger.FormatTitle("WRONG"), err)
	}

	return items
}
