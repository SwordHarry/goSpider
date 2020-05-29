package main

import (
	"../crawler/engine"
	"../crawler/gushiwen/parser"
	"../crawler/scheduler"
	"./config"
	itemSaver "./persist/client"
	"./rpcsupport"
	workerClient "./worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

const gushiwen = "https://www.gushiwen.org/"
const cncn = "https://www.cncn.com/meishi/"
const maoyan = "https://maoyan.com/board/6"

const maoyanmovieIndex = "maoyan"

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host, format: ':0000'")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts(comma separated), format:':1111,:2222,:3333'")
)

func main() {
	flag.Parse()
	// itemSaver rpc
	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	// worker rpc
	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := workerClient.CreateProcessor(pool)
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		// 通过配置的形式分别注入两个 rpc 服务
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    gushiwen,
		Parser: engine.NewFuncParser(parser.ParseThemeList, config.ParseThemeList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	// 通过 消息传递 进行 顺序分发
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
