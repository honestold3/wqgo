package main

import "fmt"

func main() {

	var a1 [3]*int
	var a2 [3]*int
	a1 = a2
	fmt.Println(a1 == a2)

	b := [4]int{1,2,3}
	fmt.Println(b)

	c := [...]int{1,2,3,5:100}
	fmt.Println(c)

	type stu struct {
		name string
		age int
	}

	s1 := stu {
		name: "haha",
		age: 11,
	}

	s2 := stu {
		"hehe",
		22,
	}
	fmt.Println(s1 == s2)

	d := [...]stu {
		{name:"hello",age:11},
		{"hahah",77},
	}
	fmt.Println(d)

	e := [2][4]int {
		{1,2,3},
	}
	fmt.Println(e)

	fmt.Println(len(d))
	fmt.Println(len(e))
	fmt.Println(cap(e))


	//数组指针， 指针数组
	x,y := 10,20
	z := [...]*int{&x,&y}
	p := &z
	fmt.Println(&z)
	fmt.Println(&p)
	fmt.Println(*p)
	fmt.Println(&x,&y)

	f := [2]int{11,22}
	var g [2]int
	g = f
	fmt.Printf("f: %p, %v \n",&f,f)
	fmt.Printf("g: %p, %v \n",&g,g)
	fmt.Println("-----------------")

	fmt.Printf("f: %p, %v \n",&f,f)
	test(&f)
	fmt.Println(f)
}

func test(x *[2]int) {
	fmt.Printf("f: %p, %v \n",&x,x)
	x[0] = 1111
}
