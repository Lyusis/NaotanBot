package utils

import (
	"logger"
	"math/rand"
	"time"
)

func Delay(waitingSeed int) {
	if waitingSeed == 0 {
		waitingSeed = 2
	}
	waitTime := waitingSeed/2 + rand.Intn(waitingSeed)
	rateLimiter := time.Tick(
		time.Duration(waitTime) * time.Second)

	logger.Sugar.Debug("Waiting", logger.FormatTitle("Time"), waitTime)

	<-rateLimiter
}
