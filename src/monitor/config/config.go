package config

import (
	"io/ioutil"
	"monitor/logger"
	"strconv"
	"strings"
)

var (
	// Wait Rate limiting seed, Second
	Wait = 1
	// RoomList Map of LivingRoom
	RoomList = make(map[int]string)
	// RoomStatusList Map of Room Status
	RoomStatusList = make(map[int]bool)
)

func init() {

	fileBytes, err := ioutil.ReadFile("src/monitor/config/virtual_liver_list.txt")
	if err != nil {
		logger.Logger.Sugar().Panic(err)
		logger.Logger.Panic("文件读取失败, 请确认后重试。")
	}
	lines := strings.Split(string(fileBytes), "\n")
	for _, line := range lines {
		kv := strings.Fields(line)
		id, idOk := strconv.Atoi(kv[0])

		if idOk != nil {
			logger.Logger.Warn("配置失败, 请确认后重试。")
		} else {
			RoomList[id] = kv[1]
		}
	}

	for index := range RoomList {
		RoomStatusList[index] = false
	}
}
