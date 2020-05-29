package client

import (
	"../../../crawler/engine"
	"../../config"
	"../../worker"
	"net/rpc"
)

// rpc worker 的链接客户端
func CreateProcessor(clientChan chan *rpc.Client) engine.RequestProcessor {

	return func(r engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializedRequest(r)

		var sResult worker.ParseResult
		client := <-clientChan
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializedResult(sResult), nil
	}
}
