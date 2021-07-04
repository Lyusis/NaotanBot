package bilibili

import (
	"sync"

	"github.com/Lyusis/NaotanMonitor/conf"
)

type LiveRoom struct {
	Name   string
	Status bool
	RoomId int
}

var (
	// RoomList Map of LivingRoom
	RoomList = make([]LiveRoom, 0)
	// mu 读写锁
	mu sync.RWMutex
)

func init() {
	for _, liver := range conf.LiverList {
		room := LiveRoom{RoomId: liver.RoomId, Name: liver.Nickname, Status: false}
		RoomList = append(RoomList, room)
	}
}

func WriteRoomStatusList(roomId int, status bool) {
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < len(RoomList); i++ {
		if RoomList[i].RoomId == roomId {
			roomP := &RoomList[i]
			roomP.Status = status
		}
	}
}

func GetRoomStatus(roomId int) bool {
	mu.RLock()
	defer mu.RUnlock()

	status := false
	for _, room := range RoomList {
		if room.RoomId == roomId {
			status = room.Status
		}
	}
	return status
}

func GetRoomName(roomId int) string {
	mu.RLock()
	defer mu.RUnlock()

	name := ""
	for _, room := range RoomList {
		if room.RoomId == roomId {
			name = room.Name
		}
	}
	return name
}
