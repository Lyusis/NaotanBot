package utils

import (
	"github.com/Lyusis/NaotanBot/conf"
	"math/rand"
	"time"

	"github.com/Lyusis/NaotanBot/logger"
)

const (
	LongDelay = 1800
)

func Delay(waitingSeed int) <-chan time.Time {
	if waitingSeed == 0 {
		waitingSeed = 2
	}
	waitTime := waitingSeed/2 + rand.Intn(waitingSeed)
	rateLimiter := time.Tick(
		time.Duration(waitTime) * time.Second)

	logger.Sugar.Debug(logger.FormatMsg("Waiting"), logger.FormatTitle("Time"), waitTime)

	if waitingSeed == LongDelay {
		logger.Sugar.Debug(logger.FormatMsg("Long Waiting Over"), logger.FormatTitle("Time"), LongDelay)
		conf.SetConf()
	}

	return rateLimiter
}
