package persist

import (
	"monitor/logger"
)

func ItemSaver() chan interface{} {
	in := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-in
			//log.Info(`Saving #itemCount #item`)
			logger.Info("Saving... \tNo.%d|\tContext: %+v", itemCount, item)
			itemCount++
		}
	}()
	return in
}
