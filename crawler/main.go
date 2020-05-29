package main

import (
	"../business/gushiwen/parser"
	"../common/engine"
	"../common/scheduler"
	"../crawler_distributed/config"
	"./persist"
)

const gushiwen = "https://www.gushiwen.org/"
const cncn = "https://www.cncn.com/meishi/"
const zhenai = "http://www.zhenai.com/zhenghun/"
const maoyan = "https://maoyan.com/board/6"

const foodSpiderIndex = "food_around"
const datingProfile = "dating_profile"
const maoyanmovieIndex = "maoyan"

func main() {
	itemChan, err := persist.ItemSaver(foodSpiderIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker, // 单机版worker
	}
	e.Run(engine.Request{
		Url:    gushiwen,
		Parser: engine.NewFuncParser(parser.ParseThemeList, config.ParseThemeList),
	})
}
