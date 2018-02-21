package main

import (
	"fmt"
	"strconv"
	//"sync"
)

type Stu struct {
	name string
	age  int
}

type N int

func (s *Stu) setAge(age int) int {
	s.age = age
	return s.age
}

func (n *N) toString() string {
	//return fmt.Sprintf("%d",n)
	return strconv.Itoa(int(*n))
}

type Stu1 struct {
	age int
	people
}

type people struct {
	name string
}

func (s Stu1) value() {
	s.age = 11
}

func (s *Stu1) pointer() {
	s.age = 11
}

func (p *people) getname() string {
	return p.name
}

func main() {

	var s Stu
	fmt.Println(s.setAge(11))

	var a N = 23
	fmt.Println(a.toString())

	fmt.Println("---------------------")

	var s1 Stu1
	s1.age = 22
	s1.value()
	fmt.Println(s1.age)
	s1.pointer()

	var s2 *Stu1
	s2 = &Stu1{
		age:    33,
		people: people{name: "hehe"},
	}
	s2.pointer()
	fmt.Println(s2.age)
	s2.value()
	fmt.Println(s2.getname())

	/*
		d := data{}
		d.Lock()
		//ye wu chu li
		d.Unlock()
	*/

	var n1 N
	var t Ner = &n1
	t.a()

}

type Ner interface {
	a()
	b(int)
	c(string) string
}

func (N) a() {

}

func (*N) b(int) {

}
func (*N) c(string) string {
	return ""
}
