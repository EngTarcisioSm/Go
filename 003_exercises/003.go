/*
- Using the solution from the previous exercise:
     1. In package-level scope, assign the following values to variables:
         1. for "x" assign 42
         2. for "y" assign "James Bond"
         3. for "z" set it to true
     2. In the main function:
         1. Use fmt.Sprintf to assign all of these values to a single variable. Do this string assignment to a variable named "s" using the short declaration operator.
         2. Demonstrate the variable "s".
*/
package main

import "fmt"

var x int 
var y string 
var z bool 

func main(){
	x := 42
	y := "James Bond"
	z := true

	s := fmt.Sprint(x, " ", y, " ", z, "\n")
    s2 := fmt.Sprintf("%v %v %v\n", x, y, z)

	fmt.Print(s)
    fmt.Print(s2)

}