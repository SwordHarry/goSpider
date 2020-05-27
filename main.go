package main

import (
	"./cncn/parser"
	"./engine"
	"./persist"
	"./scheduler"
)

const cncn = "https://www.cncn.com/meishi/"
const zhenai = "http://www.zhenai.com/zhenghun/"

const foodSpiderIndex = "food_around"
const datingProfile = "dating_profile"

func main() {
	itemChan, err := persist.ItemSaver(foodSpiderIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
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
