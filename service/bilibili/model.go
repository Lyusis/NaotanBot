package bilibili

import (
	"config"
	"sync"
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
	for _, liver := range config.LiverList {
		room := LiveRoom{RoomId: liver.RoomId, Name: liver.Name, Status: false}
		RoomList = append(RoomList, room)
	}
}

func WriteRoomStatusList(roomId int, status bool) {
	mu.Lock()
	for i := 0; i < len(RoomList); i++ {
		if RoomList[i].RoomId == roomId {
			roomP := &RoomList[i]
			roomP.Status = status
		}
	}
	mu.Unlock()
}

func GetRoomStatus(roomId int) bool {
	status := false
	mu.RLock()
	for _, room := range RoomList {
		if room.RoomId == roomId {
			status = room.Status
		}
	}
	mu.RUnlock()
	return status
}

func GetRoomName(roomId int) string {
	name := ""
	mu.RLock()
	for _, room := range RoomList {
		if room.RoomId == roomId {
			name = room.Name
		}
	}
	mu.RUnlock()
	return name
}
