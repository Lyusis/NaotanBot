package basic

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/service/cq"
)

func SendUpdateMsg() {
	if "" != conf.Announcement {
		cq.SendTool.SendGroupMessage(conf.GroupId, conf.Announcement)
	}
}

func SendQuitMsg() {
	if conf.QuitMessage != "" {
		cq.SendTool.SendGroupMessage(conf.GroupId, conf.QuitMessage)
	}
}

func Menu(msgMessage cq.MessageMessage) {
	// 命令格式
	commands := "菜单"
	// AT检测
	msgMessage.SingleAtFilter(commands, func(SendTool cq.Sender) {
		cq.SendTool.SendGroupMessage(conf.GroupId,
			"1. 骂阿骏 {words}\n2. 订阅 {uid} {nickname}\n3. 删除订阅 {keyword}\n4. 订阅列表\n5. 天气\n6. 新闻\n7. 舔狗日记")
	})
}
