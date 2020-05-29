package persist

import (
	"../engine"
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
)

func Save(client *elastic.Client, item engine.Item, index string) error {

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
