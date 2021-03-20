/*
- Using the solution from the previous exercise:
     1. In package-level scope, using the var keyword, create a variable with the identifier "y". The type of this variable must be the underlying type of the type you created in the previous exercise.
     2. In the main function:
         1. This should already be done:
             1. Demonstrate the value of the variable "x"
             2. Demonstrate the type of the variable "x"
             3. Assign 42 to the variable "x" using the operator "="
             4. Demonstrate the value of the variable "x"
         2. Now also do:
             1. Use conversion to transform the value type of the variable "x" into its underlying type and, using the "=" operator, assign the value from "x" to "y"
             2. Demonstrate the value of "y"
             3. Demonstrate the "y" type
*/


package main

import "fmt"

type jaboticaba int

var x jaboticaba

var y int

func main(){

	fmt.Printf("Value x: %v, Type x: %T\n", x, x)
	x := jaboticaba(42)
	fmt.Printf("Value x: %v, Type x: %T\n", x, x)
	y = int(x)
	fmt.Printf("Value y: %v, Type y: %T\n", y, y)
}