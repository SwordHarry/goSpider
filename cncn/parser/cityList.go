package parser

// parser: cityList -> city -> food -> store: 职责链模式

import (
	"../../engine"
	"fmt"
	"log"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a[^>]+href="(http[s]*://[a-z]+\.cncn\.com/meishi/)"[^>]+>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(matches))
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		city := string(m[2])
		log.Printf("City: %s", city)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				// cncn.com 中有根据相对定位进行url请求的： /meishi/:foodid 的情况，还要将 city 一层一层传下去
				return ParseCity(bytes, url, city)
			},
		})
		break
	}
	return result
}
