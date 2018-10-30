/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/24
 * Time: 15:07
 */
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// for循环
func forControl() {
	sum := 0
	// i := 0 初始化语句 可选的
	// i < 10 条件表达式
	// i++ 后置语句 可选的
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 去除 初始化语句 后置语句
	n := 1
	for ; n < 1000; {
		n += n
	}
	fmt.Println(n)

	// 等价于while
	m := 1
	for m < 1000 {
		m += m
	}
	fmt.Println(m)

	// 无限循环
	/*	for  {
			// do something
		}
	*/
}

var x = rand.New(rand.NewSource(time.Now().Unix())).Intn(100) - 50

func ifControl() {
	if x > 0 {
		fmt.Println("正数", x)
	} else if x == 0 {
		fmt.Println("0", x)
	} else {
		fmt.Println("负数", x)
	}
}

func switchControl() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	switch {
	case x > 0:
		fmt.Println("正数", x)
	case x == 0:
		fmt.Println("0", x)
	default:
		fmt.Println("负数", x)
	}
}

func deferTest() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	defer fmt.Println("defer最后输出")
	forControl()
	ifControl()
	switchControl()
	deferTest()
}
