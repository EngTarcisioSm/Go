package main

import "fmt"

var x int32
var z int32

func main(){
	x := 1
	z := 3000

	x = x << 1
	fmt.Printf("%b %v\n",x,x)
	z = z >> 2
	fmt.Printf("%b %v\n",z,z)

}