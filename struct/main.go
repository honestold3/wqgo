package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	name string
	age  int
}

type attr struct {
	owner int
	perm  int
}

type file struct {
	name string
	attr //组合关系
}

var kankan struct {
	name string
	age  int
}

func main() {

	u := user{
		name: "zhangsan",
		age:  22,
	}
	fmt.Println(u)

	u2 := user{"lisi", 33}
	fmt.Println(u2)

	u3 := struct {
		name string
		age  int
	}{
		name: "haha",
		age:  55,
	}
	fmt.Println(u3)

	var a struct{} //空结构，用在gorouine通讯
	var b [100]struct{}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))

	f2 := file{
		name: "a.txt",
		attr: attr{
			owner: 1,
			perm:  2,
		},
	}
	fmt.Println(f2)

	type file2 struct {
		name string
		attr struct {
			owner int
			perm  int
		}
	}

	f := file2{
		name: "aa.txt",
	}
	f.attr.owner = 1
	f.attr.perm = 2

	fmt.Println(f)

	pp := &user{
		name: "zhangsan",
		age:  22,
	}

	fmt.Println("pp:", &pp)
	fmt.Println("&pp.age:::", &pp.age)
	hehe(pp)
	fmt.Println("########################-")

	kankan.age = 11
	kankan.name = "haha"

	fmt.Println(kankan)
	fmt.Println(&kankan.name, &kankan.age)
	fmt.Println("-----------------------")
	haha()
}

func haha() {
	a := kankan
	fmt.Println(kankan)
	fmt.Println(&kankan.name, &kankan.age)

	fmt.Println(a.name, a.age)
	fmt.Println(&a.name, &a.age)

}

func hehe(u *user) {
	fmt.Println("u:::", &u)
	fmt.Println("&u.age:::", &u.age)
}
