package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name string
	age  int
}

func main() {
	x, _ := strconv.Atoi("12")
	fmt.Println(x)

}
