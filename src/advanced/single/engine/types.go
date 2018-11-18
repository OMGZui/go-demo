/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/11/18
 * Time: 15:23
 */
package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request     // 请求url
	Items    []interface{} // 城市名字
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
