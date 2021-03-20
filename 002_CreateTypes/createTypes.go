package main

import "fmt"

//created an underlying type
type hotdog int

var b hotdog = 10

func main() {

	x := 5
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("x: %v, %T\n", x, x)
	x = int(b)	//conversion
	fmt.Printf("x: %v, %T\n", x, x)

}