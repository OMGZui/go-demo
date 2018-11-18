/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/30
 * Time: 22:18
 */
package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("错误的StatusOK: %d", res.StatusCode)
	}

	bodyReader := bufio.NewReader(res.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("失败: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
