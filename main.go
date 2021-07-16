package main

import (
	"fmt"
	"github.com/Lyusis/NaotanBot/sender"
	"github.com/Lyusis/NaotanBot/service/bilibili"
	"github.com/Lyusis/NaotanBot/service/common"
)

func main() {

	fmt.Println("更新信息发送至主群")
	common.SendUpdateMsg()

	fmt.Println("CQ监听服务启动中")
	sender.HttpCQServer()
	//sender.WSCQServer()

	fmt.Println("推送服务启动中")

	fmt.Println("启动bilibili 直播通知")
	bilibili.SendLiveStatusService()

	select {}
}
