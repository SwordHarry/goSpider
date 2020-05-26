package parser

import (
	"../../engine"
	"../../model"
	"regexp"
	"strconv"
)

var storeNameRe = regexp.MustCompile(`<div class="title"><h1>(.*?)</h1>`)
var foodImgRe = regexp.MustCompile(`<div class="produce_info"><dl><dt><a><img src="(.*?)"[^>]+></a></dt><dd>`)
var costRe = regexp.MustCompile(`<dt>人&emsp;&emsp;均：</dt><dd><span class="c_f00">&yen;(\d+?)</span>`)
var phoneRe = regexp.MustCompile(`<dt>电&emsp;&emsp;话：</dt><dd><span class="c_008fe9">(.+?)</span>`)
var timeRe = regexp.MustCompile(`<dt>营业时间：</dt><dd>(.+?)</dd>`)
var deviceRe = regexp.MustCompile(`<dt>设施服务：</dt><dd>(.+?)</dd>`)
var recommendRe = regexp.MustCompile(`<dt>推&ensp;荐&ensp;菜：</dt><dd>(.+?)</dd>`)
var addressRe = regexp.MustCompile(`<dt>地&emsp;&emsp;址：</dt><dd>(.+?)&emsp;`)

func ParseStore(contents []byte, foodName string, cityName string) engine.ParseResult {
	store := model.Store{}
	store.FoodName = foodName
	store.Name = extractString(contents, storeNameRe)
	store.ImgUrl = extractString(contents, foodImgRe)
	store.Cost = extractInt(contents, costRe)
	store.Phone = extractString(contents, phoneRe)
	store.Time = extractString(contents, timeRe)
	store.Device = extractString(contents, deviceRe)
	store.Recommend = extractString(contents, recommendRe)
	store.Address = extractString(contents, addressRe)
	result := engine.ParseResult{
		Items: []interface{}{
			store,
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	val, err := strconv.Atoi(extractString(contents, re))
	if err != nil {
		return val
	}
	return val
}
