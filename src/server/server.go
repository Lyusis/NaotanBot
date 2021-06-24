package server

import (
	"api"
	"monitor/logger"
	"net/http"
)

type PostType struct {
	PostType string `json:"post_type"`
}

type MessageMessage struct {
	Time        int    `json:"time"`
	SelfId      int    `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type omitempty"`
	MessageId   int    `json:"message_id"`
	GroupId     int    `json:"group_id omitempty"`
	UserId      int    `json:"user_id"`
	TargetId    string `json:"target_id omitempty"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message omitempty"`
	Sender      struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
	} `json:"sender omitempty"`
}

type MetaEventMessage struct {
	Interval      int    `json:"interval"`
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfId        int64  `json:"self_id"`
	Status        struct {
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
		Stat           struct {
			PacketReceived  int `json:"packet_received"`
			PacketSent      int `json:"packet_sent"`
			PacketLost      int `json:"packet_lost"`
			MessageReceived int `json:"message_received"`
			MessageSent     int `json:"message_sent"`
			DisconnectTimes int `json:"disconnect_times"`
			LostTimes       int `json:"lost_times"`
			LastMessageTime int `json:"last_message_time"`
		} `json:"stat"`
	} `json:"status"`
	Time int `json:"time"`
}

func NewServer(addr string) {
	http.HandleFunc("/", handlerFunc)

	logger.Sugar.Info("启动监听服务器", logger.FormatTitle("IP地址"), addr, logger.FormatTitle("端口"), 9000)

	addr += ":9000"

	serverErr := http.ListenAndServe(addr, nil)
	if serverErr != nil {
		logger.Sugar.Error("监听启动失败", logger.FormatTitle("WRONG"), serverErr)
	}
}

func handlerFunc(_ http.ResponseWriter, r *http.Request) {
	api.CQHandler(r)
}
