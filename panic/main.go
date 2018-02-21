package main

import (
	"fmt"

	"errors"
)


func test(f func(s string) (int,error)) int {
	a, _ := f("abc")
	return a
}

func test2(s string) (int, error){

	fmt.Println(s)
	return 2,errors.New("error")
}

func testdefer() int {
	a := 100
	defer func(){
		fmt.Println("defer",a)
		a+=100
	}()
	//temp:=a
	//return temp
	return a
}

func main() {
	aa,bb := "ab","cd"
	cc := string(aa)
	fmt.Println(aa+bb+cc)

	x,y := 1,2
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ccccc",err)
		}
	}()

	fmt.Println(x, y+1)
	fmt.Println(testdefer())

	panic("is not!!!")
	fmt.Println("test panic")
}
