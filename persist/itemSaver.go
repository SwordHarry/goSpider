package persist

import (
	"../engine"
	"context"
	"github.com/pkg/errors"
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
			err := save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
