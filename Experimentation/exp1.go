package main

import "fmt"


func main(){
	x := 55
	totalB := 0
	n, _ := fmt.Println("Your age is", x)
	totalB += n
	n = 0
	totalB += foo()
	fmt.Println("Total Bytes Written:", totalB)

}

func foo() int {
	n, _ := fmt.Println("...Maybe..")
	return n
}
