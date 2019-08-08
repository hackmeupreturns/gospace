package main

import "fmt"

var a int

type maths int

var b maths

func main() {
	a = 7
	b = 8
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// can't just a=b because they are different type,
	// but we can use conversion:
	// b = maths(a) or...
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
