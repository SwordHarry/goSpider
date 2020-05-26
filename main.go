package main

import (
	"./cncn/parser"
	"./engine"
	"./persist"
	"./scheduler"
)

const cncn = "https://www.cncn.com/meishi/"
const zhenai = "http://www.zhenai.com/zhenghun/"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        cncn,
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:        zhenai,
	//	ParserFunc: parser.ParseCityList,
	//})
}
