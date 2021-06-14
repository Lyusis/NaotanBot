package utils

import (
	"monitor/config"
	"monitor/logger"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func Delay(url string) {

	log := logger.Logger{}.InitLogger().Logger
	waitTime := config.Wait + rand.Intn(config.Wait/6)
	rateLimiter := time.Tick(
		time.Duration(waitTime)*time.Second +
			time.Duration(waitTime)*time.Millisecond)

	log.Info("Fetching: ",
		zap.String("URL", url),
		zap.Int("WaitingTime", waitTime))

	<-rateLimiter

}
