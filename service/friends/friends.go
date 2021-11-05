package friends

import (
	"container/list"
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/scheduler/instance"
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/redis"
	"math/rand"
	"net/http"
)

// AJun 回复阿骏
func AJun(msgMessage cq.MessageMessage) {
	if conf.AJun == msgMessage.UserId {
		data, err := redis.SetGet("AJW")
		if err != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to get AJun Words"), err)
			cq.SendTool.SendGroupMessage(conf.GroupId, "数据异常, 无法辱骂阿骏")
		} else if len(data) != 0 {
			p := rand.Intn(len(data) - 1)
			cq.SendTool.SendGroupMessage(conf.GroupId, data[p])
		}
	}
}

// InsertAJW 添加阿骏语录
func InsertAJW(msgMessage cq.MessageMessage) {
	commands := "骂阿骏 {words}"
	msgMessage.AtFilter(commands, func(params *list.List, SendTool cq.Sender) (string, error) {
		return "辱骂语句登记异常, 无法新增语录",
			redis.SetAdd("AJW", params.Front().Value.(string))
	})
}

// InitiativeAJun 手动骂阿骏
func InitiativeAJun(msgMessage cq.MessageMessage) {
	// 命令格式
	commands := "骂阿骏"
	// AT检测
	msgMessage.SingleAtFilter(commands, func(SendTool cq.Sender) {
		data, err := redis.SetGet("AJW")
		if err != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to get AJun Words"), err)
			cq.SendTool.SendGroupMessage(conf.GroupId, "数据异常, 无法辱骂阿骏")
			return
		}
		if data != nil {
			p := rand.Intn(len(data))
			cq.SendTool.SendGroupMessage(conf.GroupId, data[p])
		}
	})
}

// Tiangou 舔狗日记
func Tiangou(msgMessage cq.MessageMessage) {
	commands := "舔狗日记"
	msgMessage.SingleAtFilter(commands, func(SendTool cq.Sender) {
		urlStr := "https://cloud.qqshabi.cn/api/tiangou/api.php"
		instance.ConcurrentEngineWorker.RequestChan <- engine.Request{
			Url:    urlStr,
			Method: http.MethodGet,
			Name:   "今日新闻",
			Parser: getNews,
		}
	})
}

func getNews(contents []byte) engine.ResultItems {

	var (
		response  string
		saveItems engine.ResultItems
	)
	response = string(contents)
	saveItems.Items = append(saveItems.Items, response)
	cq.SendTool.SendGroupMessage(conf.GroupId, response)

	return saveItems
}
