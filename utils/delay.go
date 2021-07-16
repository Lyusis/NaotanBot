package utils

import (
	"math/rand"
	"time"

	"github.com/Lyusis/NaotanBot/logger"
)

func Delay(waitingSeed int) {
	if waitingSeed == 0 {
		waitingSeed = 2
	}
	waitTime := waitingSeed/2 + rand.Intn(waitingSeed)
	rateLimiter := time.Tick(
		time.Duration(waitTime) * time.Second)

	logger.Sugar.Debug(logger.FormatMsg("Waiting"), logger.FormatTitle("Time"), waitTime)

	<-rateLimiter
}
