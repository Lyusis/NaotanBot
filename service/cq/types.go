package cq

type PostType struct {
	PostType string `json:"post_type"`
}

type MessageMessage struct {
	Sender struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
	} `json:"server omitempty"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type omitempty"`
	TargetId    string `json:"target_id omitempty"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message omitempty"`
	Time        int    `json:"time"`
	SelfId      int    `json:"self_id"`
	MessageId   int    `json:"message_id"`
	GroupId     int    `json:"group_id omitempty"`
	UserId      int    `json:"user_id"`
}

type MetaEventMessage struct {
	Status struct {
		Stat struct {
			PacketReceived  int `json:"packet_received"`
			PacketSent      int `json:"packet_sent"`
			PacketLost      int `json:"packet_lost"`
			MessageReceived int `json:"message_received"`
			MessageSent     int `json:"message_sent"`
			DisconnectTimes int `json:"disconnect_times"`
			LostTimes       int `json:"lost_times"`
			LastMessageTime int `json:"last_message_time"`
		} `json:"stat"`
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
	} `json:"status"`
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfId        int64  `json:"self_id"`
	Interval      int    `json:"interval"`
	Time          int    `json:"time"`
}

type IndividualMessage struct {
	Action string `json:"action"`
	Params struct {
		UserId  int    `json:"user_id"`
		Message string `json:"message"`
	} `json:"params"`
	Echo string `json:"echo"`
}

type GroupMessage struct {
	Action string `json:"action"`
	Params struct {
		GroupId int    `json:"group_id"`
		Message string `json:"message"`
	} `json:"params"`
	Echo string `json:"echo"`
}
