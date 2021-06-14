package model

// LiveData 直播API的response下的data数据/**
type LiveData struct {
	RoomId      int  `json:"room_id"`
	ShortId     int  `json:"short_id"`
	Uid         int  `json:"uid"`
	NeedP2p     int  `json:"need_p2p"`
	IsHidden    bool `json:"is_hidden"`
	IsLocked    bool `json:"is_locked"`
	IsPortrait  bool `json:"is_portrait"`
	LiveStatus  int  `json:"live_status"`
	HiddenTill  int  `json:"hidden_till"`
	LockTill    int  `json:"lock_till"`
	Encrypted   bool `json:"encrypted"`
	PwdVerified bool `json:"pwd_verified"`
	LiveTime    int  `json:"live_time"`
	RoomShield  int  `json:"room_shield"`
	IsSp        int  `json:"is_sp"`
	SpecialType int  `json:"special_type"`
}

// LiveResponse 直播API的response数据/**
type LiveResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Message string   `json:"message"`
	Data    LiveData `json:"data"`
}
