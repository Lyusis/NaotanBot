package utils

import (
	"math/rand"
	"monitor/logger"
	"time"
)

func Delay(waitingSeed int) {
	if waitingSeed == 0 {
		waitingSeed = 2
	}
	waitTime := waitingSeed/2 + rand.Intn(waitingSeed)
	rateLimiter := time.Tick(
		time.Duration(waitTime) * time.Second)

	logger.Sugar.Info("Waiting", logger.FormatTitle("Time"), waitingSeed)

	<-rateLimiter
}
