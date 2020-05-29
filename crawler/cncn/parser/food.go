package parser

import (
	"../../../crawler_distributed/config"
	"../../engine"
	"log"
	"regexp"
)

// 非贪婪匹配，加问号，否则只会找到最后一项
var titleRe = regexp.MustCompile(`<div class="title"><h1>(.*?)</h1>`)
var foodRe = regexp.MustCompile(`<li><a href="(.*?)"><img src="(.*?)"[^/>]*?/><div class="text_con"><strong>(.*?)</strong><p>[^<]*?</p></div></a></li>`)

func parseFood(contents []byte, cityName string, _ string) engine.ParseResult {
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
			Url:    url,
			Parser: NewStoreParser(string(title), cityName),
		})
	}
	return result
}

type FoodParser struct {
	cityName string
}

func (f *FoodParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseFood(contents, f.cityName, url)
}

func (f *FoodParser) Serialize() (name string, args interface{}) {
	return config.ParseFood, f.cityName
}

func NewFoodParser(cityName string) *FoodParser {
	return &FoodParser{cityName}
}
