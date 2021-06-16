package utils

import (
	"go.uber.org/zap"
	"monitor/config"
	"monitor/logger"
	"time"
	"math/rand"
)

func Delay(waitingSeed int) {
	
	if waitingSeed == 0 {
		waitingSeed = config.Wait
	} 
	waitTime := waitingSeed / 2 + rand.Intn(waitingSeed) 
	rateLimiter := time.Tick(
		time.Duration(waitTime) * time.Second)

	logger.Logger.Info("Waiting...", zap.Int("Time", waitingSeed))

	<-rateLimiter
}
