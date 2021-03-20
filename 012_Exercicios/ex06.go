/*
	Utilizando iota, crie 4 constantes valores sejam os proximos 4 anos.
	Demonstre pos valores 
*/

package main

import "fmt"

const(
	a = iota + 2021
	b
	c
	d
)

func main() {
	fmt.Printf("%d\n%d\n%d\n%d\n", a, b, c, d)
}