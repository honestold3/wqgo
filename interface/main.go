package main

import (
	"fmt"
)

type tester interface {
	test()
	string() string
}

type data struct {}

func (d data) test()  {
	fmt.Println("test")
}

func (d data) string() string  {
	fmt.Println("string")
	return "string"
}

func main() {

	var d data
	var i tester
	i = d
	i.test()
	i.string()

}
