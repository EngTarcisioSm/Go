package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
)

//Another way to use iota, defining only the first value as iota and the others followed, if you want to jump just use underscore
const (
	a = iota
	b
	c
	d
	_
	f
)

func main(){
	fmt.Printf("x: %v %T\ny: %v %T\nz: %v %T\n", x, x, y, y, z, z)

	fmt.Print(a,b,c,d,f)
}