package parser

import (
	"../../../common/engine"
	"log"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		log.Printf("cityList: %s", string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 用于序列化和反序列化，字符串参数为序列化json的结果
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
