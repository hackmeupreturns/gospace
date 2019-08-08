package main

import "fmt"

var x2 int
var y2 string
var z2 bool
var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true

type testInt int

var x4 testInt
var y5 int

func main() {
	fmt.Println("Exercise 1\n----------------START----------------")
	exercise1()
	fmt.Println("---------------- END ----------------")
	fmt.Println("Exercise 2\n----------------START----------------")
	exercise2()
	fmt.Println("---------------- END ----------------")
	fmt.Println("Exercise 3\n----------------START----------------")
	exercise3()
	fmt.Println("---------------- END ----------------")
	fmt.Println("Exercise 4\n----------------START----------------")
	exercise4()
	fmt.Println("---------------- END ----------------")
	fmt.Println("Exercise 5\n----------------START----------------")
	exercise5()
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

func exercise3() {
	s := fmt.Sprintf("%v\n%v\n%v", x3, y3, z3)
	fmt.Println(s)
}

func exercise4() {
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)
	x4 = 42
	fmt.Println(x4)
}

func exercise5() {
	// copy of 4:
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)
	x4 = 42
	fmt.Println(x4)
	// start of extension:
	y5 = int(x4)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
}
