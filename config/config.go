package config

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
	//WorkerCount Count of worker
	WorkerCount = 2
	// Wait Rate limiting seed, Second, default 10s
	Wait = 10
)

func init() {

	config := YAMLParsing("config.yml")
	LiverList = config.Liver
	Wait = config.Wait
	SaberchanCode = config.SaberchanCode
	GroupId = config.GroupId
	CQServer = config.CQServer
	Token = config.Token
	WorkerCount = config.WorkerCount
	QQ = config.QQ

	// RoomList[6775697] = "海苹果小学校"
	// RoomList[22470204] = "瑞芙"
	// RoomList[21672023] = "弥希"
}
