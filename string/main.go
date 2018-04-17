package main

import "fmt"

func main() {
	str := "Hello world"
	//str[0] = 'x' // 编译错误
	//如果想修改字符串的内容,可以通过如下的方式修改:
	c :=[]rune(str)
	c[0]= 'x'

	fmt.Println(string(c))

	fmt.Println("-------------------------")

	s:="hello你好"
	fmt.Println(len(s))//输出长度为11
	fmt.Println(len([]rune(s)))//输出长度为7
	s="你好"
	fmt.Println(len(s))//输出长度为6
	fmt.Println(len([]rune(s)))//输出长度为2
	s="好"
	fmt.Println([]byte(s))//输出长度为6
	fmt.Println(rune('好'))//输出20320
}
