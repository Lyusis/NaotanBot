package conf

import (
	"bufio"
	"fmt"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// 全局配置变量
var (
	// Livers vup信息
	Livers []Liver
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
	// QuitMessage Quit Message
	QuitMessage = ""
	//WorkerCount Count of worker
	WorkerCount = 0
	// Waiting Rate limiting seed, Second, default 10s
	Waiting = 0
	// ReLoad 更新通知
	ReLoad bool
)

// 本地变量
var (
	// fileName config文件地址
	fileName = "./conf/config.toml"
)

func init() {

	viper.SetDefault("worker_count", 1)
	viper.SetDefault("waiting", 60)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")
	if configReadErr := viper.ReadInConfig(); configReadErr != nil {
		logger.Sugar.Error(logger.FormatMsg("Fatal error config file"), logger.FormatError(configReadErr))
	}

	SetConf()

	viper.WatchConfig()
	watch := func(e fsnotify.Event) {
		SetConf()
		logger.Sugar.Info(logger.FormatMsg("Config file has been changed"), e.String())
		Reloading()
	}
	viper.OnConfigChange(watch)

	// RoomList[6775697] = "海苹果小学校"
	// RoomList[22470204] = "瑞芙"
	// RoomList[21672023] = "弥希"
}

func Reloading()     { ReLoad = true }
func CheckedReload() { ReLoad = false }

// SetConf 给全局变量赋值
func SetConf() {
	config := &Configuration{}
	if err := viper.Unmarshal(config); err != nil {
		logger.Sugar.Error(logger.FormatMsg("Failed to unmarshal configure file"), logger.FormatError(err))
	}
	Livers = config.Livers
	CQReceiver = config.CQReceiver
	CQSendDest = config.CQSendDest
	SaberchanCode = config.SaberchanCode
	GroupId = config.GroupId
	QQ = config.QQ
	Token = config.Token
	Announcement = config.Announcement
	QuitMessage = config.Quit
	WorkerCount = config.WorkerCount
	Waiting = config.Waiting
}

// AddListConfig 添加多层级的Toml属性
func AddListConfig(str string, sth []map[string]interface{}) error {
	var (
		mapList = make([]map[string]interface{}, 0)
	)
	switch str {
	case LiversToml:
		addLivers(&mapList, sth)
	}
	viper.Set(str, mapList)
	err := writeInto()
	return err
}

func DeleteListConfig(str string, sth []map[string]interface{}) error {
	var (
		mapList = make([]map[string]interface{}, 0)
	)
	switch str {
	case LiversToml:
		deleteLivers(&mapList, sth)
	}
	viper.Set(str, mapList)
	err := writeInto()
	return err
}

func addLivers(mapList *[]map[string]interface{}, sth []map[string]interface{}) {
	for _, liver := range Livers {
		liverMap := make(map[string]interface{})
		liverMap[NicknameToml] = liver.Nickname
		liverMap[RoomIdToml] = liver.RoomId
		*mapList = append(*mapList, liverMap)
	}
	addConfUtil(RoomIdToml, sth, mapList)
}

func deleteLivers(mapList *[]map[string]interface{}, sth []map[string]interface{}) {
	for _, liver := range Livers {
		liverMap := make(map[string]interface{})
		liverMap[NicknameToml] = liver.Nickname
		liverMap[RoomIdToml] = liver.RoomId
		*mapList = append(*mapList, liverMap)
	}
	addConfUtil(RoomIdToml, sth, mapList)
}

func AddPairConfig(str string, sth map[string]interface{}) {
	switch str {
	case CQSendDestToml:
		viper.Set("cqsenddest.ip", CQSendDest.IP)
		viper.Set("cqsenddest.port", CQSendDest.Port)
	case CQReceiverToml:
		viper.Set("cqreceiver.ip", CQReceiver.IP)
		viper.Set("cqreceiver.port", CQReceiver.Port)
	default:
		logger.Sugar.Error(logger.FormatMsg("Failed to write new config"), logger.FormatError(fmt.Errorf("cannot find attribute")))
	}
	for key, value := range sth {
		name := str + "." + key
		viper.Set(name, value)
	}
	writeInto()
}

func AddSimpleConfig(str string, sth interface{}) {
	viper.Set(str, sth)
	writeInto()
}

// mapList 最终输出值, 基于原本的数据 A map
// sth 新增值, B map
func addConfUtil(keyAttribute string, sth []map[string]interface{}, mapList *[]map[string]interface{}) {
	sthTransform := transformMap(keyAttribute, sth)
	mapListTransform := transformMap(keyAttribute, *mapList)
	mergeMapTransform := coverAndMergeMap(mapListTransform, sthTransform)
	mergeMap := unTransformMap(keyAttribute, mergeMapTransform)
	*mapList = mergeMap
}
func transformMap(keyAttribute string, mapList []map[string]interface{}) map[interface{}]map[string]interface{} {
	transform := make(map[interface{}]map[string]interface{})
	for _, object := range mapList {
		keyValue := object[keyAttribute]
		transform[keyValue] = object
	}
	return transform
}
func unTransformMap(keyAttribute string, transform map[interface{}]map[string]interface{}) []map[string]interface{} {
	mapList := make([]map[string]interface{}, 0)
	for _, value := range transform {
		mapList = append(mapList, value)
	}
	return mapList
}

// coverAndMergeMap B map 覆盖或合并到A map上
// 当 A map 有但 B map 没有的值, 保持原样添加到新 map
// 当 B map 有但 A map 没有的值, 添加到新 map
// 当 A B map 都有的值, 将 B map 的值添加到新 map
func coverAndMergeMap(aMap, bMap map[interface{}]map[string]interface{}) map[interface{}]map[string]interface{} {
	mergeMap := aMap
	for key, value := range bMap {
		mergeMap[key] = value
	}
	return mergeMap
}

func writeInto() error {
	err := viper.WriteConfigAs(fileName)
	if err != nil {
		logger.Sugar.Error(logger.FormatMsg("Failed to write new config"), logger.FormatError(err))
	}
	return err
}

func DeleteConfig(key, value string) {
	file, openFileErr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if openFileErr != nil {
		logger.Sugar.Error(logger.FormatMsg("Failed to open config file"), logger.FormatError(openFileErr))
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Sugar.Error(logger.FormatMsg("Failed to close config file"), logger.FormatError(err))
		}
	}(file)

	reader := bufio.NewReader(file)
	for {
		line, readStringErr := reader.ReadString('\n')
		if readStringErr != nil {
			logger.Sugar.Error(logger.FormatMsg("Failed to read config file"), logger.FormatError(readStringErr))
			return
		}
		if strings.Contains(line, key) {

		}
	}
}
