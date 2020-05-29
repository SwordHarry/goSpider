package parser

// parser: cityList -> city -> food -> store: 职责链模式

import (
	"../../engine"
	"../../model"
	"log"
	"regexp"
	"strings"
)

var movieRe = regexp.MustCompile(`<a href="(/films/\d+)" title="(.*?)" class="image-link" data-act="boarditem-click" data-val="{movieId:(\d+)}">`)
var coverRe = regexp.MustCompile(`<img data-src="(.*?)"[^c]*class="board-img"[^/]*/>`)

//var actorRe = regexp.MustCompile(`<p class="star">主演：(.*?)</p>`) // 有的没主演
var timeRe = regexp.MustCompile(`<p class="releasetime">上映时间：(.*?)</p>`)
var nextRe = regexp.MustCompile(`<a class="page_\d+?"[^h]*?href="(.*?)"[^>]*?>下一页</a>`)

func ParseMovie(contents []byte, url string) engine.ParseResult {
	matches := movieRe.FindAllSubmatch(contents, -1)
	covers := coverRe.FindAllSubmatch(contents, -1)
	//actorList := actorRe.FindAllSubmatch(contents, -1)
	timeList := timeRe.FindAllSubmatch(contents, -1)
	log.Printf("total num of movie in this page: %d", len(matches))
	result := engine.ParseResult{
		Requests: []engine.Request{},
		Items:    []engine.Item{},
	}
	for i, m := range matches {
		url := "maoyan.com" + string(m[1])
		title := string(m[2])
		id := string(m[3])
		coverUrl := string(covers[i][1])
		//actors := string(actorList[i][1])
		time := string(timeList[i][1])
		movie := model.Movie{
			Name:     title,
			CoverUrl: coverUrl,
			//Actors:   actors,
			Time: time,
		}
		//fmt.Println(url, id, movie)
		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Id:      id,
			Type:    "movie",
			Payload: movie,
		})
	}

	nextPage := nextRe.FindSubmatch(contents)
	if len(nextPage) > 0 {
		index := strings.Index(url, "?")
		if index != -1 {
			url = url[:index]
		}
		// 下一页
		result.Requests = append(result.Requests, engine.Request{
			Url:    url + string(nextPage[1]),
			Parser: engine.NewFuncParser(ParseMovie, "ParseMovie"),
		})
	}

	return result
}
