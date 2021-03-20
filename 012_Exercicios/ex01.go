package main

import "fmt"

var i int

func main(){
	i := 32

	fmt.Printf("%d - %b - %#x\n", i, i, i)
}

/*
	para expressar um numero em decimais  é possivel  utilizar %v ou %d
*/