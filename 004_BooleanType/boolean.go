package main

import "fmt"

var x bool

func main(){
	fmt.Println(x)
	x = true 
	fmt.Println(x)
	x = (10 < 100)
	fmt.Println(x)

	y := (100 == 10)
	z := (10 > 100)
	fmt.Println(y)
	fmt.Println(z)
}
