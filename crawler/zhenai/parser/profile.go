package parser

import (
	"../../engine"
	"../../model"
	"regexp"
	"strconv"
)

// 标题
//var nickNameRe = regexp.MustCompile(`<h1 class="nickName" [^>]*>([^<]+)</h1>`)

// 紫色按钮
// profileRe 取第一个和最后一个匹配项为 婚姻 和 教育
var purpleProfileRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)岁</div>`)
var constellationRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>(.+)座([^<])</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>(\d+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>(\d+)kg</div>`)
var workPlaceRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>工作地:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([^<]+)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>(.*\\.*)</div>`)

// 粉色按钮
//var pinkProfileRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>(.*房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>(.*车)</div>`)
var huKouRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>籍贯:([^<]+)</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name

	// 页面上紫色按钮为个人重要信息，但是有的人没有列出体重，导致数量不一样
	profileMatch := purpleProfileRe.FindAllSubmatch(contents, -1)[0]
	profile.Marriage = string(profileMatch[1])
	profile.Education = string(profileMatch[len(profileMatch)-1])
	extractInt(contents, ageRe, func(val int) {
		profile.Age = val
	})
	profile.Constellation = extractString(contents, constellationRe)
	extractInt(contents, heightRe, func(val int) {
		profile.Height = val
	})
	extractInt(contents, weightRe, func(val int) {
		profile.Weight = val
	})
	profile.WorkPlace = extractString(contents, workPlaceRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Occupation = extractString(contents, occupationRe)

	// 页面上粉红色按钮为附属信息
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	profile.Hukou = extractString(contents, huKouRe)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
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

func extractInt(contents []byte, re *regexp.Regexp, callback func(val int)) {
	val, err := strconv.Atoi(extractString(contents, re))
	if err != nil {
		callback(val)
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		name,
	}
}
