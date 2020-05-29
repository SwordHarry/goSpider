package engine

import (
	"log"
)

type ConcurrentEngine struct {
	WorkerCount int
	ItemChan    chan Item
	Scheduler
	RequestProcessor
}
type RequestProcessor func(r Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	// WorkerChan: 向 scheduler 询问 channel，可能是 simple 版的共用channel，也可能是队列版的每个独立channel
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

var visitedUrlMap = make(map[string]bool)

// 去重
func isDuplicate(url string) bool {
	if visitedUrlMap[url] {
		return true
	}
	visitedUrlMap[url] = true
	return false
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("# %v Got item: %v", itemCount, item)
			itemCount++
			// ATTENTION: 闭包的坑
			go func(curItem Item) { e.ItemChan <- curItem }(item)
		}

		for _, request := range result.Requests {
			// 一次去重
			if isDuplicate(request.Url) {
				continue
			}
			e.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			// worker: fetcher + parser
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
