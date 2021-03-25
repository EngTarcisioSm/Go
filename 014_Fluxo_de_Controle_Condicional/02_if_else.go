/*
	- if, else							- se, então 
	- if, else if, else					- se, então se, entao
	- if, else if, ..., else if, else

*/
package main

import "fmt"

func main(){
	x := 10

	if x > 100{
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}

	y := 25

	if y < 30{
		fmt.Println("3")
	} else if y > 50 {		//pode haver quantos iff else forem necessários 
		fmt.Println("4")
	} else {
		fmt.Println("5")
	}
}