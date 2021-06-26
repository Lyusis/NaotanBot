package config

var (
	// RoomList Map of LivingRoom
	RoomList = make(map[int]string)
	// RoomStatusList Map of Room Status
	RoomStatusList = make(map[int]bool)
	// SaberchanCode Saberchan send code
	SaberchanCode = ""
	// GroupId Id of QQ Group
	GroupId = ""
	// CQServer Ip of go-cqhttp Server
	CQServer = ""
	//WorkerCount Count of worker
	WorkerCount = 2
	// Wait Rate limiting seed, Second, default 10s
	Wait = 10
)

func init() {

	config := YAMLParsing("src/config/config.yml")
	livers := config.Items
	Wait = config.Wait
	SaberchanCode = config.SaberchanCode
	GroupId = config.GroupId
	CQServer = config.CQServer
	WorkerCount = config.WorkerCount

	for _, liver := range livers {
		RoomList[liver.RoomId] = liver.Name
	}

	// RoomList[6775697] = "海苹果小学校"
	// RoomList[22470204] = "瑞芙"
	// RoomList[21672023] = "弥希"

	for index := range RoomList {
		RoomStatusList[index] = false
	}
}
