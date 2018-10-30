/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/25
 * Time: 11:10
 */
package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

type Vertex struct {
	X float64
	Y float64
}

// Abs 方法拥有一个名为 v，类型为 Vertex 的接收者
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 普通函数
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func fn1() {
	// 方法就是一类带特殊的 接收者 参数的函数
	v1 := Vertex{3, 4}
	v2 := Vertex{3, 4}
	fmt.Println(v1.Abs())

	// 指针
	v1.Scale(10) // 等价 	(&v1).Scale(10)
	fmt.Println(v1.Abs())

	// 普通函数传值
	ScaleFunc(&v2, 10)
	fmt.Println(AbsFunc(v2))
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func fn2() {
	var i I = T{"hello"}
	//i := T{"oo"}
	i.M()

	// 空接口
	var f interface{}
	describe(f)

	f = 42
	describe(f)

	f = "oo"
	describe(f)

	// 类型断言
	var h interface{} = "hello"
	s, ok := h.(string)
	fmt.Println(s, ok)
}

func describe(i interface{}) {
	fmt.Printf("%v %T\n", i, i)
}

// 类型选择
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("整型：%v\n", v)
	case string:
		fmt.Printf("字符串：%q\n", v)
	default:
		fmt.Printf("布尔：%v\n", v)
	}
}

func fn3() {
	do(21)
	do("hello")
	do(true)

	// Stringer
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a)
	fmt.Println(z)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func fn4() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Reader
	r := strings.NewReader("123456789")
	b := make([]byte, 4)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:%v] = %q\n", n, b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	fn1()
	fn2()
	fn3()
	fn4()
}
