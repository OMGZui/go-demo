/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/11/18
 * Time: 15:36
 */
package engine

import (
	"github.com/omgzui/go-demo/src/advanced/single/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("FetchUrl: %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("失败,Url %s : %v", r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Item: %v", item)
		}
	}
}
