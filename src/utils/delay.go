package utils

import (
	"config"
	"math/rand"
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

	logger.Info("Waiting", true, "Time", waitingSeed)

	<-rateLimiter
}
