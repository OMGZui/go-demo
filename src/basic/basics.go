/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/24
 * Time: 14:26
 */
package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool
var php, perl = "php", "perl"

const PI = 3.14

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("我最喜欢的数字是：", rand.Intn(10))
	fmt.Println(math.Pi)

	fmt.Println("1+1 =", add(1, 1))

	x, y := swap(1, 2)
	fmt.Println("1 2交换数字：", x, y)

	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(php, perl)
	k := 3
	fmt.Println(k)

	// go 基本类型
	// bool
	// string
	// int uint
	// byte uint8别名
	// rune int32别名
	// float32 float64
	// complex64 complex128
	fmt.Printf("Type:%T, Value:%v\n", k, k)
	fmt.Println(PI)
}
