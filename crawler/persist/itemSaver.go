package persist

import (
	"../../common/engine"
	"../../common/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// 客户端维护集群，内网看不见， turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item %d: %v", itemCount, item)
			itemCount++
			err := persist.Save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}
