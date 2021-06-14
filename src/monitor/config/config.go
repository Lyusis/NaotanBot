package config

import (
	"monitor/logger"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	// Wait Rate limiting seed, Second
	Wait = 60
	// RoomList Map of LivingRoom
	RoomList = make(map[int]string)
)

func Init() {
	log := logger.Logger{}.InitLogger().Logger

	fileBytes, err := ioutil.ReadFile("src/monitor/config/virtual_liver_list.txt")
	if err != nil {
		log.Sugar().Panic(err)
		log.Panic("文件读取失败, 请确认后重试。")
	}
	lines := strings.Split(string(fileBytes), "\n")
	for _, line := range lines {
		kv := strings.Fields(line)
		id, idOk := strconv.Atoi(kv[0])

		if idOk != nil {
			log.Warn("配置失败, 请确认后重试。")
		} else {
			RoomList[id] = kv[1]
		}
	}
}
