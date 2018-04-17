package main


/*
#include <stdlib.h>

typedef int (*intFunc) ();

int bridge_int_func(intFunc f)
{
	return f();
}

int fortytwo()
{
    return 42;
}
*/
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {

	Seed(22)
	fmt.Println("kankan:",Random())

	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))

}
