package parser

import (
	"../../../crawler_distributed/config"
	"../../engine"
	"../../model"
	"log"
	"regexp"
)

var sentenceRe = regexp.MustCompile(`<div class="cont"[^>]*?>[\s\S]*?<a[^h]*?href="https://so.gushiwen.cn/mingju/juv_([0-9a-zA-Z]*?).aspx">([^<]*?)</a>`)
var fromRe = regexp.MustCompile(`<span[^>]*?>——</span><a[^>]*?>([^<]*?)</a>`)

func parseSentence(contents []byte, url string, theme string) engine.ParseResult {
	sentenceMatches := sentenceRe.FindAllSubmatch(contents, -1)
	fromMatches := fromRe.FindAllSubmatch(contents, -1)
	log.Printf("get sentence len: %v", len(sentenceMatches))
	result := engine.ParseResult{}
	for i, s := range sentenceMatches {
		id := string(s[1])
		content := string(s[2])
		from := string(fromMatches[i][1])
		sentence := model.Sentence{
			Theme:   theme,
			Content: content,
			From:    from,
		}
		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Id:      id,
			Type:    "sentence",
			Payload: sentence,
		})
	}
	return result
}

type SentenceParser struct {
	theme string
}

func (s *SentenceParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseSentence(contents, url, s.theme)
}

func (s *SentenceParser) Serialize() (name string, args interface{}) {
	return config.ParseSentence, s.theme
}

func NewSentenceParser(theme string) *SentenceParser {
	return &SentenceParser{theme: theme}
}
