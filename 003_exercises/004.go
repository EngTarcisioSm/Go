/*
- Create a type. The underlying type must be int.
- Create a variable for this type, with the identifier "x", using the keyword var.
- In the main function:
     1. Demonstrate the value of the variable "x"
     2. Demonstrate the type of the variable "x"
     3. Assign 42 to the variable "x" using the operator "="
     4. Demonstrate the value of the variable "x"
*/
package main

import "fmt"

type jaboticaba int

var x jaboticaba

func main(){

	fmt.Printf("Value: %v, Type: %T\n", x, x)
	x := 42
	fmt.Printf("Value: %v, Type: %T\n", x, x)
}