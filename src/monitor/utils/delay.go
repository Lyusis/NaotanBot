package utils

import (
	"math/rand"
	"monitor/config"
	"monitor/logger"
	"time"
)

func Delay(waitingSeed int) {
	if waitingSeed == 0 {
		waitingSeed = config.Wait
	}
	waitTime := waitingSeed/2 + rand.Intn(waitingSeed)
	rateLimiter := time.Tick(
		time.Duration(waitTime) * time.Second)

	logger.Info("Waiting...\t| Time: %d", waitingSeed)

	<-rateLimiter
}
