/*
- Use var to declare three variables. They must have pacake-level scope. Do not assign values to these variables. Use the following identifiers and types for these variables:
1. x must have type int
2. y must have type string
3. z must be of type bool

- In the main function
1. Demonstrate the values of each identifier
2. The compiler assigned values for these variables. What are these values called?
*/
package main

import "fmt"

var x int 
var y string 
var z bool 

func main(){

	fmt.Printf("Type: %T", x)
	fmt.Printf("Type: %T", y)
	fmt.Printf("Type: %T", z)

	fmt.Println("all variables have nil values")

	fmt.Printf("Value x: %v\n", x)
	fmt.Printf("Value y: %v\n", y)
	fmt.Printf("Value c: %v\n", z)
}