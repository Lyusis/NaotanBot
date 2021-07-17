package main

import (
	"fmt"
	"github.com/Lyusis/NaotanBot/server"
	"github.com/Lyusis/NaotanBot/service/bilibili"
)

func main() {

	fmt.Println("更新信息发送至主群")
	//common.SendUpdateMsg()

	fmt.Println("CQ监听服务启动中")
	//server.HttpCQServer()
	server.WSCQServer()

	fmt.Println("推送服务启动中")

	fmt.Println("启动bilibili 直播通知")
	bilibili.SendLiveStatusService()

	select {}
}
