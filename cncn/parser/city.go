package parser

import (
	"../../engine"
	"fmt"
	"log"
	"regexp"
)

// 非贪婪匹配，加问号，否则只会找到最后一项
var cityRe = regexp.MustCompile(`<a href="/meishi/([0-9]+/)" target="_blank">.*?<div><span>(.*?)</span>.*?</p></div></a>`)

func ParseCity(contents []byte, urlPrefix string, cityName string) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(matches))
	result := engine.ParseResult{}

	log.Printf("total num of Food: %d", len(matches))
	for _, m := range matches {
		name := string(m[2])
		url := urlPrefix + string(m[1])
		log.Printf("Food: %s", name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: foodParser(cityName),
		})
	}
	return result
}

func foodParser(cityName string) engine.ParserFunc {
	return func(contents []byte, url string) engine.ParseResult {
		return ParseFood(contents, cityName, url)
	}
}
