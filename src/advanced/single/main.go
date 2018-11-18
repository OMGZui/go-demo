/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/30
 * Time: 20:19
 */
package main

import (
	"github.com/omgzui/go-demo/src/advanced/single/engine"
	"github.com/omgzui/go-demo/src/advanced/single/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
