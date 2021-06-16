package config

import (
	"io/ioutil"
	"monitor/logger"
	"strconv"
	"strings"
)

var (
	// Wait Rate limiting seed, Second
	Wait = 10
	// RoomList Map of LivingRoom
	RoomList = make(map[int]string)
	// RoomStatusList Map of Room Status
	RoomStatusList = make(map[int]bool)
)

func init() {

	fileBytes, err := ioutil.ReadFile("src/monitor/config/virtual_liver_list.txt")
	if err != nil {
		logger.Panic("文件读取失败\t%+v", err)
	}
	lines := strings.Split(string(fileBytes), "\n")
	for _, line := range lines {
		kv := strings.Fields(line)
		id, idErr := strconv.Atoi(kv[0])

		if idErr != nil {
			logger.Panic("配置失败\t%+v", idErr)
		} else {
			RoomList[id] = kv[1]
		}
	}

	for index := range RoomList {
		RoomStatusList[index] = false
	}
}
