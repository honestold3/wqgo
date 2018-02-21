package main

import (
	"fmt"
	"reflect"
)

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct {}

type stu struct {
	name string
	age int
}

func (d *data) test(){}

func (d *data) string() string {
	return "aaa"
}

func pp(a stringer) {
	fmt.Println(a.string())
}

func main() {

	var d data
	var t tester = &d
	t.test()
	fmt.Println(t.string())

	pp(t)
	var s stringer = t
	pp(s)

	t1 := 11
	var f float32 = float32(t1)
	fmt.Println(f)

	var i interface{}
	i = 4
	fmt.Println(i)
	i = "abc"
	fmt.Println(i)

	m := make(map[string]interface{},0)
	m["key1"] = 11
	m["key2"] = "value2"

	s1 := stu{
		name: "haha",
		age: 32,
	}
	m["key3"] = s1
	fmt.Println(m)

	d1 := m["key1"].(int)//类型断言
	fmt.Println(d1)

	var d2 int
	d2 = d1
	fmt.Println(d2)
	fmt.Println(reflect.TypeOf(d1))

	var s2 stu
	s2,ok := m["key3"].(stu)
	fmt.Println(s1.name,s2.age,ok)

	d3,ok := m["key4"].(int)
	fmt.Println(ok,d3)

}