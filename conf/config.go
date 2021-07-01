package conf

import (
	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	LiverList []Liver
	// SaberchanCode Saberchan send code
	SaberchanCode = ""
	// QQ Bot's QQ
	QQ = ""
	// GroupId Id of QQ Group
	GroupId = ""
	// CQServer Ip of go-cqhttp Server
	CQServer = ""
	// Token Access token
	Token = ""
	// Announcement Update Message
	Announcement = ""
	//WorkerCount Count of worker
	WorkerCount = 0
	// Waiting Rate limiting seed, Second, default 10s
	Waiting = 0
)

func init() {

	viper.SetDefault("worker_count", 1)
	viper.SetDefault("waiting", 60)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")
	if configReadErr := viper.ReadInConfig(); configReadErr != nil {
		if _, ok := configReadErr.(viper.ConfigFileNotFoundError); ok {
			logger.Sugar.Error(logger.FormatMsg("Cannot find the configure file"), logger.FormatTitle("WRONG"), configReadErr)
		} else {
			logger.Sugar.Error(logger.FormatMsg("Unable to read the configure file"), logger.FormatTitle("WRONG"), configReadErr)
		}
	}

	SetConf()

	// Config Watching
	viper.WatchConfig()
	watch := func(e fsnotify.Event) {
		logger.Sugar.Info(logger.FormatMsg("Config file is changed: %s \n"), e.String())
		SetConf()
	}
	viper.OnConfigChange(watch)

	// RoomList[6775697] = "海苹果小学校"
	// RoomList[22470204] = "瑞芙"
	// RoomList[21672023] = "弥希"
}

func SetConf() {
	config := &Configuration{}
	if err := viper.Unmarshal(config); err != nil {
		logger.Sugar.Error(logger.FormatMsg("Failed to unmarshal configure file"), logger.FormatTitle("WRONG"), err)
	}
	LiverList = config.Livers
	Waiting = config.Waiting
	WorkerCount = config.WorkerCount
	SaberchanCode = config.SaberchanCode
	GroupId = config.GroupId
	CQServer = config.CQServer
	QQ = config.QQ
	Token = config.Token
	Announcement = config.Announcement
}
