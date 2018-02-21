package main

import "fmt"

type data struct {
	x int
	s string
}

func main() {

	const (
		read byte = 1 << iota
		write
		exec
		freeze
	)

	fmt.Printf("read: %b,write: %b,exec: %b,freeze: %b\n", read, write, exec, freeze)

	fmt.Println("--------------------------------")
	a := read | write | freeze
	b := read | write | exec
	c := a & b
	fmt.Printf("%04b & %04b = %04b\n", a, b, c)
	fmt.Println("--------------------------------")

	d := 1
	p := &d
	*p++
	fmt.Println(*p)

	x := 10
	var p1 *int = &x
	*p += 20
	fmt.Println(p1, *p1)

	p2 := &x
	fmt.Println(p2)
	p2 = nil
	fmt.Println(p2)

	fmt.Println("--------------------------------")

	var a1 data = data{1, "abc"}
	fmt.Println(a1)
	b1 := data{
		2,
		"abcd",
	}

	c1 := &data{
		x: 3,
		s: "hjk",
	}

	fmt.Println(b1, c1)
	fmt.Println("--------------------------------")

	n := 0
	r := &n
	fmt.Println("r:", *r)
	play(10, 5, r)
	fmt.Println("play:", *r)
}

func play(a int, b int, r *int) {
	*r = a * b
}
