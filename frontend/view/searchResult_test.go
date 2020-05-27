package view

import (
	"../../engine"
	itemModel "../../model"
	"../model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	item := engine.Item{
		Url:  "",
		Type: "cncn",
		Id:   "",
		Payload: itemModel.Store{
			Name:      "全聚德",
			FoodName:  "北京烤鸭",
			CityName:  "北京",
			ImgUrl:    "",
			Cost:      80,
			Phone:     "1234566",
			Time:      "8-9",
			Device:    "免费wifi",
			Address:   "北京市",
			Recommend: "",
		},
	}
	page := model.SearchResult{
		Hits:     10,
		Start:    0,
		Items:    []engine.Item{},
		Query:    "",
		PrevFrom: 0,
		NextFrom: 0,
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	out, err := os.Create("template.test.html")
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
