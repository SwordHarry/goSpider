package parser

import (
	"../../../common/engine"
	"../../../crawler_distributed/config"
	"fmt"
	"log"
	"regexp"
)

// 非贪婪匹配，加问号，否则只会找到最后一项
var cityRe = regexp.MustCompile(`<a href="/meishi/([0-9]+/)" target="_blank">.*?<div><span>(.*?)</span>.*?</p></div></a>`)

func parseCity(contents []byte, urlPrefix string, cityName string) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(matches))
	result := engine.ParseResult{}

	log.Printf("total num of Food: %d", len(matches))
	for _, m := range matches {
		name := string(m[2])
		url := urlPrefix + string(m[1])
		log.Printf("Food: %s", name)
		result.Requests = append(result.Requests, engine.Request{
			Url:    url,
			Parser: NewFoodParser(cityName),
		})
	}
	return result
}

type CityParser struct {
	cityName string
}

func (c *CityParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseCity(contents, url, c.cityName)
}

func (c *CityParser) Serialize() (name string, args interface{}) {
	return config.ParseCity, c.cityName
}

func NewCityParser(cityName string) *CityParser {
	return &CityParser{cityName: cityName}
}
