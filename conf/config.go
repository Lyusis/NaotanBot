package conf

import (
	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	LiverList []Liver
	// CQReceiver Ip of receive information
	CQReceiver = Addr{}
	// CQSendDest IP of send information
	CQSendDest = Addr{}
	// SaberchanCode Saberchan send code
	SaberchanCode = ""
	// QQ Bot's QQ
	QQ = ""
	// GroupId Id of QQ Group
	GroupId = ""
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
			logger.Sugar.Error(logger.FormatMsg("Cannot find the configure file"), logger.FormatError(configReadErr))
		} else {
			logger.Sugar.Error(logger.FormatMsg("Unable to read the configure file"), logger.FormatError(configReadErr))
		}
	}

	SetConf()

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

// SetConf 给全局变量赋值
func SetConf() {
	config := &Configuration{}
	if err := viper.Unmarshal(config); err != nil {
		logger.Sugar.Error(logger.FormatMsg("Failed to unmarshal configure file"), logger.FormatError(err))
	}
	LiverList = config.Livers
	CQReceiver = config.CQReceiver
	CQSendDest = config.CQSendDest
	SaberchanCode = config.SaberchanCode
	GroupId = config.GroupId
	QQ = config.QQ
	Token = config.Token
	Announcement = config.Announcement
	WorkerCount = config.WorkerCount
	Waiting = config.Waiting
}
