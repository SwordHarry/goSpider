package parser

import (
	"../../engine"
	"fmt"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a[^>]+href="(http[s]*://[a-z]+\.cncn\.com/meishi/)"[^>]+>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(matches))
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				// cncn.com 中有根据相对定位进行url请求的： /meishi/:foodid 的情况
				return ParseCity(bytes, url)
			},
		})
	}
	return result
}
