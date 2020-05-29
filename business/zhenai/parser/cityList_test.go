package parser

import (
	"../../../common/model"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)

	result := ParseCityList(contents, "")
	// verify result
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		hukou := result.Items[i].Payload.(model.Profile).Hukou
		if hukou != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, hukou)
		}
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items; but had %d", resultSize, len(result.Items))
	}
}
