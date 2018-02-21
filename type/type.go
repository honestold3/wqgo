package main

import (
	"fmt"
	"math"
)

type Student struct {
	Age  int
	Name string
}

func main() {
	/*
		bool
		byte
		int 8 16 32 64 int8 int16
		uint
		float8 16 32 64
		complex64 128 1+2i
		uintptr
		string
		array
		struct
		function
		interface
		map
		slice
		channel
	*/

	s := "0Abc"
	fmt.Println([]byte(s))
	fmt.Println("---------------------------")

	stu := new(Student)
	stu.Age = 11
	stu.Name = "wang"
	fmt.Println(*stu)
	fmt.Println("---------------------------")

	var kk *Student = &Student{22, "rrr"}
	fmt.Println(*kk)
	fmt.Println("---------------------------")

	stu1 := &Student{
		33,
		"zhang",
	}
	fmt.Println(*stu1)
	fmt.Println("---------------------------")

	fmt.Println(math.MinInt8, math.MaxUint16)
	fmt.Println("---------------------------")

	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	fmt.Println(m)
	fmt.Println("---------------------------")

	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)
	fmt.Println("---------------------------")

	p := new(map[int]string)
	m1 := *p
	m[1] = "1"
	fmt.Println(m1)
	fmt.Println("---------------------------")

	a := 1
	b := uint8(a)
	fmt.Println(b)
	fmt.Println("---------------------------")

	type color int

	var red color
	red = 1
	var green color
	green = 2
	fmt.Println(red, green)

	red1 := int(1)
	//red1 = red
	fmt.Println(red1)

	//reflect.Kind

}
