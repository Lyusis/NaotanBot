package bilibili

import (
	"sync"

	"github.com/Lyusis/NaotanBot/conf"
)

type liveRoom struct {
	Liver  conf.Liver
	Status bool
}

var (
	// RoomList Map of LivingRoom
	RoomList = make([]liveRoom, 0)
	// mu 读写锁
	mu sync.RWMutex
)

func init() {
	loadRoomList()
}

func loadRoomList() {
	RoomList = make([]liveRoom, 0)
	for _, liver := range conf.Livers {
		room := liveRoom{Liver: conf.Liver{RoomId: liver.RoomId, Nickname: liver.Nickname}, Status: false}
		RoomList = append(RoomList, room)
	}
}

func reloadRoomList() {
	oldList := RoomList
	loadRoomList()
	for _, room := range oldList {
		if room.Status {
			writeRoomStatusList(room.Liver.RoomId, true)
		}
	}
}

func writeRoomStatusList(roomId int, status bool) {
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < len(RoomList); i++ {
		if RoomList[i].Liver.RoomId == roomId {
			roomP := &RoomList[i]
			roomP.Status = status
			break
		}
	}
}

func getRoomStatus(roomId int) bool {
	mu.RLock()
	defer mu.RUnlock()

	status := false
	for _, room := range RoomList {
		if room.Liver.RoomId == roomId {
			status = room.Status
		}
	}
	return status
}

func getRoomName(roomId int) string {
	mu.RLock()
	defer mu.RUnlock()

	name := ""
	for _, room := range RoomList {
		if room.Liver.RoomId == roomId {
			name = room.Liver.Nickname
		}
	}
	return name
}
