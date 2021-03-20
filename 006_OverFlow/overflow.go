package main

import "fmt"

func main(){
	var a uint16

	// a = 65536 // overflow uint16

	a = 65535
	fmt.Println(a)
	a++
	fmt.Println(a)
	a++
	fmt.Println(a)
}