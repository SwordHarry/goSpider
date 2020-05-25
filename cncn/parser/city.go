package parser

import (
	"../../engine"
	"fmt"
	"regexp"
)

// 非贪婪匹配，加问号，否则只会找到最后一项
var cityRe = regexp.MustCompile(`<a href="/meishi/([0-9]+/)" target="_blank">.*?<div><span>(.*?)</span>.*?</p></div></a>`)

func ParseCity(contents []byte, urlPrefix string) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(matches))
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		url := urlPrefix + string(m[1])
		result.Items = append(result.Items, "food: "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseFood,
		})
	}
	return result
}
