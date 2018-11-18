/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/11/18
 * Time: 15:54
 */
package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("cityList_test_data.html")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)
	result := ParserCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("应该有: %d, 实际有：%d\n", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("应该是: %s, 实际是: %s", url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("应该有: %d"+"实际有：%d\n", resultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("应该是: %s, 实际是: %s", city, result.Items[i].(string))
		}
	}

}
