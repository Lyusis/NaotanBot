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
	// SaberchanCode Saberchan's send code
	SaberchanCode = "SCT45921Tqj6arbImzDYshqstl5siyKf9"
	// GroupId Id of QQ Group
	GroupId = "540419281"
	// CQServer Ip of go-cqhttp Server
	CQServer = "127.0.0.1"
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

	// RoomList[6775697] = "海苹果小学校"
	// RoomList[22470204] = "瑞芙"
	// RoomList[21672023] = "弥希"

	for index := range RoomList {
		RoomStatusList[index] = false
	}
}
