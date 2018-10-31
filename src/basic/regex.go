/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/30
 * Time: 21:36
 */
package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is 862275679@qq.com@a.com
aa@qq.com
bb@v.com
cc@cc.com.cn
`

func regex1() {
	//re := regexp.MustCompile("862275679@qq.com")
	//re := regexp.MustCompile(`.+@.+\..+`)
	//re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	//re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)

	m1 := re.FindString(text)
	m2 := re.FindAllString(text, -1)
	m3 := re.FindAllStringSubmatch(text, -1)
	fmt.Println(m1, m2)
	fmt.Println(m3)
}

func main() {
	regex1()
}
