package main

import "fmt"

type I interface {
	Get() int
	Put(int)
}

type S struct { i int }
func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct { i int }
func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

func main() {
	var i I
	s := &S{11}
	i = s
	f(i)

	fmt.Println("--------------------")

	r := &R{22}
	i = r
	f(i)

}

func f(p I) {
	switch t := p.(type) {
	case *S:
		fmt.Println("i store *S", t)
	case *R:
		fmt.Println("i store *R", t)
	default:
		fmt.Println("default", t)
	}
}