package parser

import (
	"../../engine"
	"fmt"
	"regexp"
)

// 非贪婪匹配，加问号，否则只会找到最后一项
var titleRe = regexp.MustCompile(`<div class="title"><h1>(.*?)</h1>`)
var foodRe = regexp.MustCompile(`<li><a href="(.*?)"><img src="(.*?)"[^/>]*?/><div class="text_con"><strong>(.*?)</strong><p>[^<]*?</p></div></a></li>`)

func ParseFood(contents []byte) engine.ParseResult {
	title := titleRe.FindSubmatch(contents)[1] // 菜名
	matches := foodRe.FindAllSubmatch(contents, -1)
	fmt.Println(string(title), len(matches))
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[3])
		url := string(m[1])
		result.Items = append(result.Items, "store: "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseStore(bytes, string(title))
			},
		})
	}
	return result
}
