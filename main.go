package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"

	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/service/basic"
	biliServer "github.com/Lyusis/NaotanBot/service/bilibili/server"
	dailiesServer "github.com/Lyusis/NaotanBot/service/dailylife/server"
	cqServer "github.com/Lyusis/NaotanBot/service/server"
	"github.com/Lyusis/NaotanBot/utils"
)

func main() {
	fmt.Println("CQ监听服务启动")
	cqServer.WSCQServer()

	fmt.Println("基础服务启动")
	basic.SendUpdateMsg()
	dailiesServer.Clock()

	fmt.Println("启动bilibili 直播通知")
	biliServer.SendLiveStatusService()

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
