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
		logger.Logger.Sugar().Panic(err)
		logger.Logger.Panic("文件读取失败, 请确认后重试。")
	}
	lines := strings.Split(string(fileBytes), "\n")
	for _, line := range lines {
		kv := strings.Fields(line)
		id, idErr := strconv.Atoi(kv[0])

		if idErr != nil {
			logger.Logger.Sugar().Panic(idErr)
			logger.Logger.Panic("配置失败, 请确认后重试。")
		} else {
			RoomList[id] = kv[1]
		}
	}

	for index := range RoomList {
		RoomStatusList[index] = false
	}
}
