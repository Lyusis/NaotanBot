package persist

import (
	"monitor/logger"
	"go.uber.org/zap"
)

func ItemSaver() chan interface{} {
	log := logger.Logger{}.InitLogger().Logger

	in := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-in
			//log.Info(`Saving #itemCount #item`)
			log.Info("Saving: ",
				zap.Int("No", itemCount),
				zap.Any("Context", item))
			itemCount++
		}
	}()
	return in
}
