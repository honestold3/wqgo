package main

import "fmt"

func main() {

	a1 := [3]string{}
	a2 := [3]string{}

	a1 = a2

	fmt.Println(a1 == a2)
	fmt.Println(a1)
	fmt.Println(a2)

	c := [...]int{1, 2, 3, 5: 100}
	fmt.Println(c)

	type stu struct {
		name string
		age  int
	}

	d := [...]stu{
		{name: "xww1", age: 11},
		{name: "xww2", age: 120},
	}
	fmt.Println(d)

	e := [...][4]int{
		{1, 2, 3},
	}
	fmt.Println(e)

	fmt.Println(len(e))
	fmt.Println(cap(e))

	x, y := 10, 20
	z := [...]*int{&x, &y}
	p := &z
	fmt.Println(*p)
	fmt.Println(&x, &y)

	f := [2]int{11, 22}
	var g [2]int
	g = f
	fmt.Printf("x: %p, %v \n", &f, f)
	fmt.Printf("x: %p, %v \n", &g, g)

	//test(&f)
	fmt.Println(f)

	var n N = 11

	fmt.Println("------------------")

	n.test()
	fmt.Println("n ", &n, n)

	n.test1()
	fmt.Println("n1 ", &n, n)

	fmt.Println(kankan())

}

type N int

func (n N) test() {
	n++
	fmt.Println("test", &n, n)
}

func (n *N) test1() {
	*n++
	fmt.Println("test1", n, *n)

}

func kankan() (int,string,int) {
	return 1,"dd",4
}

/*
func test(x *[2]int)  {
	fmt.Printf("x: %p, %v \n",&x,x)
	x[0]=1111
}
*/
