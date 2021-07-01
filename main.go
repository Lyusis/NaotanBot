package main

import (
	"fmt"
	"github.com/Lyusis/NaotanMonitor/server"
	"github.com/Lyusis/NaotanMonitor/service/bilibili"
)

func main() {

	//cq.SendQQGroupMessage(conf.GroupId, "重启中")

	fmt.Println("CQ监听服务启动中")
	server.CQServer()

	fmt.Println("推送服务启动中")

	fmt.Println("启动bilibili 直播通知")
	bilibili.SendLiveStatusService()

	select {}
}
