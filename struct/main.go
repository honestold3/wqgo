package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	name string
	age int
}

type attr struct {
	owner int
	perm int
}

type file struct {
	name string
	attr  //组合关系
}

func main() {

	u := user {
		name: "zhangsan",
		age: 22,
	}
	fmt.Println(u)

	u2 := user{"lisi",33}
	fmt.Println(u2)

	u3 := struct {
		name string
		age int
	} {
		name: "haha",
		age: 55,
	}
	fmt.Println(u3)

	var a struct{} //空结构，用在gorouine通讯
	var b [100]struct{}
	fmt.Println(unsafe.Sizeof(a),unsafe.Sizeof(b))

	f2 := file {
		name: "a.txt",
		attr: attr {
			owner: 1,
			perm: 2,
		},
	}
    fmt.Println(f2)

	type file2 struct {
		name string
		attr struct {
			owner int
			perm int
		}
	}

	f := file2{
		name: "aa.txt",
	}
	f.attr.owner = 1
	f.attr.perm = 2

	fmt.Println(f)
}
