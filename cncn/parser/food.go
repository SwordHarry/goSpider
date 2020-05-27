package parser

import (
	"../../engine"
	"log"
	"regexp"
)

// 非贪婪匹配，加问号，否则只会找到最后一项
var titleRe = regexp.MustCompile(`<div class="title"><h1>(.*?)</h1>`)
var foodRe = regexp.MustCompile(`<li><a href="(.*?)"><img src="(.*?)"[^/>]*?/><div class="text_con"><strong>(.*?)</strong><p>[^<]*?</p></div></a></li>`)

func ParseFood(contents []byte, cityName string, _ string) engine.ParseResult {
	// TODO: range out of index? matches len is 0
	title := titleRe.FindSubmatch(contents)[1] // 菜名
	matches := foodRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	log.Printf("total num of Store: %d", len(matches))
	for _, m := range matches {
		name := string(m[3])
		url := string(m[1])
		log.Printf("Store: %s", name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: StoreParser(string(title), cityName),
		})
	}
	return result
}

func StoreParser(title, cityName string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		// 放入 门店 url， 菜名， 城市名
		return ParseStore(c, url, title, cityName)
	}
}
