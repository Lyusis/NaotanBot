package main

import (
	"config"
	"fmt"
	"server"
	"service"
)

func main() {

	//cq.SendQQGroupMessage(config.GroupId, "重启中")

	fmt.Println("CQ监听服务启动中")
	go func() {
		server.NewServer(config.CQServer)
	}()

	fmt.Println("推送服务启动中")

	fmt.Println("启动bilibili 直播通知")
	service.SendLiveStatusService()

	select {}
}
