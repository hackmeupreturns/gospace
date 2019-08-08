package main

import "fmt"

var x2 int
var y2 string
var z2 bool

func main() {
	fmt.Println("Exercise 1\n----------------START----------------")
	exercise1()
	fmt.Println("---------------- END ----------------")
	fmt.Println("Exercise 2\n----------------START----------------")
	exercise2()
	fmt.Println("---------------- END ----------------")

}

func exercise1() {
	x1 := 42
	y1 := "James Bond"
	z1 := true

	fmt.Println(x1, y1, z1)
	fmt.Printf("%v\n", x1)
	fmt.Printf("%v\n", y1)
	fmt.Printf("%v\n", z1)
}

func exercise2() {
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)
	fmt.Println("The compiler assigns 0 for integers, an empty string for strings\nand false for booleans which have not been manually assigned")
}
