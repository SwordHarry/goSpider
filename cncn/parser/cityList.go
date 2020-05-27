package parser

// parser: cityList -> city -> food -> store: 职责链模式

import (
	"../../engine"
	"log"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a[^>]+href="(http[s]*://[a-z]+\.cncn\.com/meishi/)"[^>]+>([^<]+)</a>`)

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	log.Printf("total num of City: %d", len(matches))
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		city := string(m[2])
		log.Printf("City: %s", city)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: cityParser(city),
		})
	}
	return result
}

func cityParser(cityName string) engine.ParserFunc {
	// cncn.com 中有根据相对定位进行url请求的： /meishi/:foodid 的情况，还要将 city 一层一层传下去
	return func(contents []byte, url string) engine.ParseResult {
		return ParseCity(contents, url, cityName)
	}
}
