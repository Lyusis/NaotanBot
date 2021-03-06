package bilibili

// LiveDataResponse 通用response数据/**
//type LiveDataResponse struct {
//	Data    LiveDataResponseData `json:"data"`
//	Msg     string               `json:"msg"`
//	Message string               `json:"message"`
//	Status  int                  `json:"status"`
//}

type LiveDataResponseData struct {
	Title            string `json:"title"`
	RoomId           int    `json:"room_id"`
	Uid              int    `json:"uid"`
	Online           int    `json:"online"`
	LiveTime         int    `json:"live_time"`
	LiveStatus       int    `json:"live_status"`
	ShortId          int    `json:"short_id"`
	Area             int    `json:"area"`
	AreaName         string `json:"area_name"`
	AreaV2Id         int    `json:"area_v2_id"`
	AreaV2Name       string `json:"area_v2_name"`
	AreaV2ParentName string `json:"area_v2_parent_name"`
	AreaV2ParentId   int    `json:"area_v2_parent_id"`
	Uname            string `json:"uname"`
	Face             string `json:"face"`
	TagName          string `json:"tag_name"`
	Tags             string `json:"tags"`
	CoverFromUser    string `json:"cover_from_user"`
	Keyframe         string `json:"keyframe"`
	LockTill         string `json:"lock_till"`
	HiddenTill       string `json:"hidden_till"`
	BroadcastType    int    `json:"broadcast_type"`
}

type LiveDataResponse struct {
	Data    map[string]LiveDataResponseData `json:"data"`
	Msg     string                          `json:"msg"`
	Message string                          `json:"message"`
	Code    int                             `json:"code"`
}

type LivingUrl struct {
	Data    LivingUrlData `json:"data"`
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Ttl     int           `json:"ttl"`
}

type LivingUrlData struct {
	PlayurlInfo struct {
		Playurl struct {
			GQnDesc []struct {
				Desc string `json:"desc"`
				Qn   int    `json:"qn"`
			} `json:"g_qn_desc"`
			Stream []struct {
				Format []struct {
					Codec []struct {
						UrlInfo []struct {
							Host      string `json:"host"`
							Extra     string `json:"extra"`
							StreamTtl int    `json:"stream_ttl"`
						} `json:"url_info"`
						CodecName string `json:"codec_name"`
						BaseUrl   string `json:"base_url"`
						AcceptQn  []int  `json:"accept_qn"`
						CurrentQn int    `json:"current_qn"`
					} `json:"codec"`
					FormatName string `json:"format_name"`
				} `json:"format"`
				ProtocolName string `json:"protocol_name"`
			} `json:"stream"`
			P2PData struct {
				MServers interface{} `json:"m_servers"`
				P2P      bool        `json:"p2p"`
				MP2P     bool        `json:"m_p2p"`
				P2PType  int         `json:"p2p_type"`
			} `json:"p2p_data"`
			DolbyQn interface{} `json:"dolby_qn"`
			Cid     int         `json:"cid"`
		} `json:"playurl"`
		ConfJson string `json:"conf_json"`
	} `json:"playurl_info"`
	AllSpecialTypes []interface{} `json:"all_special_types"`
	IsHidden        bool          `json:"is_hidden"`
	IsLocked        bool          `json:"is_locked"`
	IsPortrait      bool          `json:"is_portrait"`
	Encrypted       bool          `json:"encrypted"`
	PwdVerified     bool          `json:"pwd_verified"`
	RoomId          int           `json:"room_id"`
	ShortId         int           `json:"short_id"`
	Uid             int           `json:"uid"`
	LiveStatus      int           `json:"live_status"`
	HiddenTill      int           `json:"hidden_till"`
	LockTill        int           `json:"lock_till"`
	LiveTime        int           `json:"live_time"`
	RoomShield      int           `json:"room_shield"`
}
