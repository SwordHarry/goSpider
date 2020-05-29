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
			Url:    url,
			Parser: NewCityParser(city),
		})
	}
	return result
}
