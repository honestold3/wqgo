package common

import "fmt"

var ss = "ss"

func sayHello2()  {
	fmt.Println("hello2")
	SayHello()
}
func init()  {
	fmt.Println("init3")
}
