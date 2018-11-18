/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/11/18
 * Time: 15:11
 */
package init

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	res, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("code：", res.StatusCode)
	}

	e := determineEncoding(res.Body)
	utf8Reader := transform.NewReader(res.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		panic(err)
	}

	printCityList(all)
	//fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

func printCityList(contents []byte) {
	//re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	//matches := re.FindAll(contents, -1)
	//for _, m := range matches {
	//	fmt.Printf("%s\n", m)
	//}
	//fmt.Printf("Found %d\n", len(matches))
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	for _, sub := range m {
	//		fmt.Printf("%s ", sub)
	//	}
	//	fmt.Println()
	//}
	for _, m := range matches {
		fmt.Printf("City: %s URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Found %d\n", len(matches))
}
