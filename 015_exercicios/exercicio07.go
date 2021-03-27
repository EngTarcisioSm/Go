/*
	crie um programa utilizando a declaração switch, sem switch statement espeficicando
*/
package main

import "fmt"

func main(){
	x:= 12

	switch{
	case x == 1:
		fmt.Print(1)
	case x == 2:
		fmt.Print(2)
	case x == 3:
		fmt.Print(3)
	case x == 4:
		fmt.Print(4)
	default:
		fmt.Print("nENHUMA DAS ANTERIORES")
	}
}