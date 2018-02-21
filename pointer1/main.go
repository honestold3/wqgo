package main

import (
	"fmt"

	"errors"
)

func check(x int) error {
	if x <= 0 {
		return errors.New("x<0")
	}
	return nil
}

type N int

func (n N) test() {
	n++
	fmt.Println("test: ", &n, n)
}

func (n *N) test1() {
	*n++
	fmt.Println("test1:", n, *n)

}

func main() {
	/*x := -3
	if err := check(x); err == nil{
		x++
		fmt.Println(x)
	} else {
		fmt.Println(err)
	}
	*/
	fmt.Println("--------------------------------")

	var n N = 11

	n.test()
	fmt.Println("n:    ", &n, n)

	n.test1()
	fmt.Println("n1:   ", &n, n)

}
