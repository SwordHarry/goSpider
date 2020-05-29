package persist

import (
	"../engine"
	"../model"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	profile := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:          "安静的雪",
			Gender:        "女",
			Age:           34,
			Height:        162,
			Weight:        57,
			Income:        "3001-5000元",
			Marriage:      "离异",
			Education:     "大学本科",
			Occupation:    "人事/行政",
			Hukou:         "山东菏泽",
			Constellation: "牡羊座",
			House:         "已购房",
			Car:           "未购车",
			WorkPlace:     "山东菏泽",
		},
	}
	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		fmt.Println("1")
		panic(err)
	}
	const index = "dating_test"
	// save expected item
	err = save(client, profile, index)
	if err != nil {
		fmt.Println("2")
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(profile.Type).Id(profile.Id).Do(context.Background())
	if err != nil {
		fmt.Println("3")
		panic(err)
	}

	t.Logf("%s", resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*(resp.Source), &actual)

	if err != nil {
		fmt.Println("4")
		panic(err)
	}
	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != profile {
		t.Errorf("got %v;expected %v", actual, profile)
	}
}
