package parser

import (
	"../../../common/engine"
	"log"
	"regexp"
)

var themeRe = regexp.MustCompile(`<a href="(https://so\.gushiwen\.org/mingju/Default\.aspx\?.*?)" target="_blank">(.*?)</a>`)

func ParseThemeList(contents []byte, url string) engine.ParseResult {
	matches := themeRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		theme := string(m[2])
		log.Printf("theme: %s", theme)

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 用于序列化和反序列化，字符串参数为序列化json的结果
			Parser: NewSentenceParser(theme),
		})
	}
	return result
}
