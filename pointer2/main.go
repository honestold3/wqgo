package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := 100
	var p *int = &a
	fmt.Println(&a)
	fmt.Printf("%p, %v\n", &p, p)
	test(p)
	test(&a)
	test(p)

	fmt.Println("--------------------")
	a1 := []int{1, 2, 3, 5}
	test1("aaa", a1[:]...)
	fmt.Println(reflect.TypeOf(a1))
	b := a1[:]
	fmt.Println(reflect.TypeOf(b))

	fmt.Println("--------------------")
	fmt.Println(test2(1))
}

func test(x *int) {
	fmt.Printf("%p, %v\n", &x, x)
}

func test1(s string, a ...int) {
	fmt.Printf("%T, %v\n", a, a)
}

func test2(a int) (int, string) {
	age := a / 3
	return age, "dd"
}
