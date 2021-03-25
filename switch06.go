/*
	pode ser feito teste(s) condicionais 
*/
package main

import "fmt"

func main(){
	for i := 1; i <= 100; i++ {
		switch {
			case (i == 1), (i > 2 && i < 10):
				fmt.Println("Opção A")
			case (i == 25), (i < 40):
				fmt.Println("Opção B")
			case i == 50:
				fmt.Println("Opção C")
			default:
				fmt.Println("Opção D")
		}
	}
}