package main

import (
	"fmt"
	"github.com/Lyusis/NaotanBot/server"

	"github.com/Lyusis/NaotanBot/service/basic"
	"github.com/Lyusis/NaotanBot/service/bilibili"
)

func main() {
	fmt.Println("CQ监听服务启动中")
	server.WSCQServer()

	fmt.Println("更新信息发送至主群")
	basic.SendUpdateMsg()

	fmt.Println("推送服务启动中")

	fmt.Println("启动bilibili 直播通知")
	bilibili.SendLiveStatusService()

	select {}
}
