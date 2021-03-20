package main

import "fmt"

/*
- Using the short declaration operator, assign these values to variables with the identifiers "x", "y", and "z".
     1. 42
     2. "James Bond"
     3. true
- Now demonstrate the values in these variables, with:
     1. A single print statement
     2. Multiple print statements
*/

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x,y,z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}