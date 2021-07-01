package main

import (
	"fmt"
	"github.com/Lyusis/NaotanMonitor/sender"
	"github.com/Lyusis/NaotanMonitor/service/bilibili"
	"github.com/Lyusis/NaotanMonitor/service/common"
)

func main() {

	fmt.Println("更新信息发送至主群")
	common.SendUpdateMsg()

	fmt.Println("CQ监听服务启动中")
	sender.CQServer()

	fmt.Println("推送服务启动中")

	fmt.Println("启动bilibili 直播通知")
	bilibili.SendLiveStatusService()

	select {}
}
