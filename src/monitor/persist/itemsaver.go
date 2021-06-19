package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"monitor/logger"
)

func ItemSaver() chan interface{} {
	in := make(chan interface{})
	go func() {
		for {
			item := <-in
			//log.Info(`Saving #itemCount #item`)
			logger.Info("Saving", false, item)
			save(item)
		}
	}()
	return in
}

func save(item interface{}) {
	client, clientErr := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if clientErr != nil {
		logger.Error("启动ElasticSearch服务失败", false, clientErr)
	}

	response, responseErr := client.Index().
		Index("dating_profile").
		BodyJson(item).
		Do(context.Background())
	if responseErr != nil {
		logger.Error("向ElasticSearch保存失败", false, responseErr)
	}
	fmt.Printf("%s\n", response.Id)
}
