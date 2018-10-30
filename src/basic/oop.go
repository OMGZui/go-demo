/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/30
 * Time: 11:17
 */
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func structOOP() {
	p := Person{"omgzui", 25}
	fmt.Println(p.name, p.age)
}

func (person *Person) setName(name string) {
	person.name = name
}

func (person *Person) getName() {
	fmt.Printf("name=>%s age=>%d\n", person.name, person.age)
}

func setGet() {
	p := new(Person)
	p.getName()
	p.setName("zzzz")
	p.getName()
}

type Student struct {
	Person
	id int
}

func extendsOOP() {
	s := Student{
		Person{"hhh", 25},
		1,
	}
	s.getInfo()
	s.setName("gggg")
	s.getInfo()
}

func (student *Student) getInfo() {
	fmt.Printf("name=>%s age=>%d id=>%d\n", student.name, student.age, student.id)
}

type Human interface {
	speak(language string)
}

type Chinese struct {
}

type American struct {
}

func (ch Chinese) speak(language string) {
	fmt.Printf("speck %s\n", language)
}

func (am American) speak(language string) {
	fmt.Printf("speck %s\n", language)
}

func oop() {
	var ch Human
	var am Human

	ch = Chinese{}
	am = American{}

	ch.speak("Chinese")
	am.speak("English")
}

func main() {
	structOOP()
	setGet()
	extendsOOP()
	oop()
}
