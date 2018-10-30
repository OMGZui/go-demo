/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/24
 * Time: 15:47
 */
package main

import (
	"fmt"
	"math"
	"strings"
)

func pointType() {
	i := 42
	p := &i
	fmt.Println("指针开始", i, p, *p)
	*p = 50
	fmt.Println("指针结束", i, p, *p)
}

type Vertex struct {
	X int
	Y int
}

func structType() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 10
	fmt.Println(v)

	// 结构体指针
	p := &v
	p.X = 20
	fmt.Println(v)

	v1 := Vertex{1, 2}  // has type Vertex
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p1 := &Vertex{1, 2} // has type *Vertex
	fmt.Println(v1, v2, v3, p1)
}

type Vertex2 struct {
	Lat, Long float64
}

func arrayType() {
	// 数组
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 切片
	// 切片并不存储任何数据，它只是描述了底层数组中的一段
	// 更改切片的元素会修改其底层数组中对应的元素
	var s = primes[1:4]
	fmt.Println(s)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Printf("%T %v\n", s2, s2)
	fmt.Println(s2[1].i)

	// make
	c1 := make([]int, 0, 5)
	c2 := make([]int, 5)
	fmt.Println(c1, c2, c1[:2], c2[:2], cap(c1), cap(c2))

	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// range
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// map
	m := make(map[string]string)
	m["a"] = "1"
	m["b"] = "2"
	fmt.Println(m, m["a"])

	m2 := map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funcType() {
	// fn 函数值
	hyPot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hyPot(5, 12))

	fmt.Println(compute(hyPot))
	fmt.Println(compute(math.Pow))

	// 闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pointType()
	structType()
	arrayType()
	funcType()
}
