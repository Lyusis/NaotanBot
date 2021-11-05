package bilibili

import (
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/redis"
	"strconv"
	"strings"
	"sync"

	"github.com/Lyusis/NaotanBot/conf"
)

type Liver struct {
	Nickname string
	Uid      int
}

type liveStatus struct {
	Liver  Liver
	Status bool
}

var (
	Reload bool
	// LiverList Map of LivingRoom
	LiverList = make([]liveStatus, 0)
	// mu 读写锁
	mu sync.RWMutex
)

func init() {
	loadRoomList()
}

func loadRoomList() {
	var (
		liversData []string
		err        error
	)
	liversData, err = redis.SetGet("Liver")
	if err != nil {
		cq.SendTool.SendGroupMessage(conf.GroupId, "获取liver信息失败")
		return
	}
	LiverList = make([]liveStatus, 0)

	for _, liver := range liversData {
		info := strings.Split(liver, ":")
		if len(info) == 2 {
			var (
				nickname string
				uid      int
			)
			uid, err = strconv.Atoi(info[0])
			if err != nil {
				break
			}
			nickname = info[1]
			room := liveStatus{Liver: Liver{Uid: uid, Nickname: nickname}, Status: false}
			LiverList = append(LiverList, room)
		}
	}
}

func reloadRoomList() {
	oldList := LiverList
	loadRoomList()
	cq.SendTool.SendGroupMessage(conf.GroupId, "正在更新Liver列表")
	for _, room := range oldList {
		if room.Status {
			writeRoomStatusList(room.Liver.Uid, true)
		}
	}
}

func writeRoomStatusList(uid int, status bool) {
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < len(LiverList); i++ {
		if LiverList[i].Liver.Uid == uid {
			roomP := &LiverList[i]
			roomP.Status = status
			break
		}
	}
}

func getRoomStatus(uid int) bool {
	mu.RLock()
	defer mu.RUnlock()

	status := false
	for _, room := range LiverList {
		if room.Liver.Uid == uid {
			status = room.Status
		}
	}
	return status
}

func getRoomName(uid int) string {
	mu.RLock()
	defer mu.RUnlock()

	name := ""
	for _, liver := range LiverList {
		if liver.Liver.Uid == uid {
			name = liver.Liver.Nickname
		}
	}
	return name
}
