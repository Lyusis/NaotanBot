package dailylife

import (
	"encoding/json"
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/scheduler/instance"
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/utils"
	"github.com/robfig/cron"
	"net/http"
	"time"
)

func GetNews(msgMessage cq.MessageMessage) {
	commands := "新闻"
	msgMessage.SingleAtFilter(commands, func(SendTool cq.Sender) {
		urlStr := "http://api.tianapi.com/bulletin/index?key=" + conf.NewsKey
		instance.ConcurrentEngineWorker.RequestChan <- engine.Request{
			Url:    urlStr,
			Method: http.MethodPost,
			Name:   "今日新闻",
			Parser: getNews,
		}
	})
}

func getNews(contents []byte) engine.ResultItems {

	var (
		response  = news{}
		newsMsg   = "近期新闻\n"
		saveItems engine.ResultItems
	)
	jsonErr := json.Unmarshal(contents, &response)
	if jsonErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { News } message"), logger.FormatError(jsonErr))
	}
	newsData := response.Newslist
	saveItems.Items = append(saveItems.Items, newsData)

	for index, newsInfo := range newsData {
		newsMsg += "\n报道时间: " + newsInfo.Mtime + "\n" +
			newsInfo.Title + "\n" + newsInfo.Digest + "\n"
		if index == 5 {
			break
		}
	}
	cq.SendTool.SendGroupMessage(conf.GroupId, newsMsg)

	return saveItems
}

func Clock() {
	go func() {
		var (
			c   = cron.New()
			err error
		)
		const (
			spec9  = "00 00 9 * * ?"
			spec21 = "00 00 21 * * ?"
			spec   = "00 00 0,1,2,3,4,5,6,7,8,10,11,12,13,14,15,16,17,18,19,20,22,23 * * ?"
		)

		err = c.AddFunc(spec9, func() {
			currentDay := time.Now()
			HelloWorld(currentDay, "早上好")
		})
		err = c.AddFunc(spec21, func() {
			currentDay := time.Now()
			HelloWorld(currentDay, "夜深了")
		})
		err = c.AddFunc(spec, func() {
			currentDay := time.Now()
			cq.SendTool.SendGroupMessage(conf.GroupId, utils.MiddleInt("现在是北京时间", currentDay.Hour(), "点整"))
		})
		if err != nil {
			cq.SendTool.SendGroupMessage(conf.GroupId, "计时任务异常")
		}
		c.Start()
		select {}
	}()
}

func HelloWorld(currentDay time.Time, hello string) {
	helloMsg := hello + "\n" +
		// 年月日星期
		"今天是" +
		utils.SingleFrontInt(currentDay.Year(), "年") +
		utils.SingleFrontInt(int(currentDay.Month()), "月") +
		utils.SingleFrontInt(currentDay.Day(), "日") +
		" 星期" + weekdayCN[currentDay.Weekday()] + "\n" +
		// 时间
		utils.MiddleInt("现在是北京时间", currentDay.Hour(), "点整")

	cq.SendTool.SendGroupMessage(conf.GroupId, helloMsg)
	urlStr := "https://api.seniverse.com/v3/weather/daily.json?key=" + conf.WeatherSecret + "&location=shanghai&unit=c&start=0&days=3"
	instance.ConcurrentEngineWorker.RequestChan <- engine.Request{
		Url:    urlStr,
		Method: http.MethodGet,
		Name:   "上海天气",
		Parser: getWeather,
	}
}

func GetWeather(msgMessage cq.MessageMessage) {
	commands := "天气"
	msgMessage.SingleAtFilter(commands, func(SendTool cq.Sender) {
		urlStr := "https://api.seniverse.com/v3/weather/daily.json?key=" + conf.WeatherSecret + "&location=shanghai&unit=c&start=0&days=3"
		instance.ConcurrentEngineWorker.RequestChan <- engine.Request{
			Url:    urlStr,
			Method: http.MethodGet,
			Name:   "上海天气",
			Parser: getWeather,
		}
	})
}

func getWeather(contents []byte) engine.ResultItems {

	var (
		response   = weathers{}
		weatherMsg = "近期的天气情况为\n"
		saveItems  engine.ResultItems
	)
	jsonErr := json.Unmarshal(contents, &response)
	if jsonErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { Weathers } message"), logger.FormatError(jsonErr))
	}
	weatherData := response.Results[0]
	saveItems.Items = append(saveItems.Items, weatherData)

	for _, day := range weatherData.Daily {
		weatherMsg += day.Date + ": " + "\n" +
			"\t白天: " + day.TextDay + " 夜晚: " + day.TextNight + "\n" +
			"\t最高温度: " + day.High + " 最低温度: " + day.Low + "\n" +
			"\t相对湿度: " + day.Humidity + "\n"
	}
	cq.SendTool.SendGroupMessage(conf.GroupId, weatherMsg)

	return saveItems
}
