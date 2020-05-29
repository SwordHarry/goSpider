package worker

// 序列化与反序列化层：把 engine 的引擎的 Request 、 ParseResult 和 rpc 层的 Request 、 ParseResult 相互转换
import (
	cncnParser "../../crawler/cncn/parser"
	"../../crawler/engine"
	gushiwenParser "../../crawler/gushiwen/parser"
	maoyanParser "../../crawler/maoyan/parser"
	"../config"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type SerializedParser struct {
	FuncName string
	Args     interface{}
}

/**
序列化结果：
maoyan:{"ParseMovie", nil}
zhenai:{"ProfileParser", userName}
*/

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

// engin.Request -> Request
func SerializedRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			name, args,
		},
	}
}

// engine.ParseResult -> ParseResult
func SerializedResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializedRequest(req))
	}
	return result
}

// Request -> engine.Request
func DeserializedRequest(r Request) (engine.Request, error) {
	parser, err := deserializedParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

// SerializedParser -> Parser
func deserializedParser(p SerializedParser) (engine.Parser, error) {
	// 服务注册与发现
	switch p.FuncName {
	// gushiwen
	case config.ParseThemeList:
		return engine.NewFuncParser(gushiwenParser.ParseThemeList, config.ParseThemeList), nil
	case config.ParseSentence:
		if theme, ok := p.Args.(string); ok {
			return gushiwenParser.NewSentenceParser(theme), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	// maoyan
	case config.ParseMovie:
		return engine.NewFuncParser(maoyanParser.ParseMovie, config.ParseMovie), nil
	// cncn
	case config.ParseCityList:
		return engine.NewFuncParser(cncnParser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		if cityName, ok := p.Args.(string); ok {
			return cncnParser.NewCityParser(cityName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	// 有参
	case config.ParseFood:
		if cityName, ok := p.Args.(string); ok {
			return cncnParser.NewFoodParser(cityName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

// ParseResult -> engine.ParseResult
func DeserializedResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializedRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}
