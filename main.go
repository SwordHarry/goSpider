package main

import (
	"./cncn/parser"
	"./engine"
	"./scheduler"
)

const seed = "https://www.cncn.com/meishi/"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        seed,
		ParserFunc: parser.ParseCityList,
	})
}
