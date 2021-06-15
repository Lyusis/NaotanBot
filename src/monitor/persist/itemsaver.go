package persist

import (
	"go.uber.org/zap"
	"monitor/logger"
)

func ItemSaver() chan interface{} {
	in := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-in
			//log.Info(`Saving #itemCount #item`)
			logger.Logger.Info("Saving: ",
				zap.Int("No", itemCount),
				zap.Any("Context", item))
			itemCount++
		}
	}()
	return in
}
