package main

import "fmt"

var c int

func main(){
	a:= 1 == 2
	b:= 2 <= 3
	c = 4
	d:= 5 != 6
	e:= 7 < 8
	f:= 9 > 10

	fmt.Print(a,b,c,d,e,f)
}