package main

import (
	"fmt"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/service/basic"
	biliServer "github.com/Lyusis/NaotanBot/service/bilibili/server"
	cqServer "github.com/Lyusis/NaotanBot/service/server"
	"github.com/Lyusis/NaotanBot/utils"
	_ "net/http/pprof"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("CQ监听服务启动中")
	cqServer.WSCQServer()

	fmt.Println("更新信息发送至主群")
	basic.SendUpdateMsg()

	fmt.Println("推送服务启动中")

	fmt.Println("启动bilibili 直播通知")
	biliServer.SendLiveStatusService()

	//select {}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	select {
	case <-quit:
		logger.Sugar.Info(logger.FormatMsg("quit bot"))
		fmt.Println("退出程序中...")
		basic.SendQuitMsg()
		<-utils.Delay(2)
		break
	}
}
