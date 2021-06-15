package utils

import (
	"go.uber.org/zap"
	"monitor/config"
	"monitor/logger"
	"time"
)

func Delay() {

	waitTime := config.Wait
	rateLimiter := time.Tick(
		time.Duration(waitTime) * time.Second)

	logger.Logger.Info("Waiting...", zap.Int("Time", waitTime))

	<-rateLimiter
}
