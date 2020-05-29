package main

import (
	"../../config"
	"../../persist"
	"../../rpcsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on itemSaver")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElastciIndexForGushiwen))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	// type ItemSaverService has no exported methods of suitable type (hint: pass a pointer to value of that type)
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
