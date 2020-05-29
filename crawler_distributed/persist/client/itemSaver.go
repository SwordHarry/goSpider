package client

import (
	"../../../crawler/engine"
	"../../config"
	"../../rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
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

			// Call rpc to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}
