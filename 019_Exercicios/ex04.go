/*
	- Crie e use um struct anônimo.
	- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
package main

import "fmt"

func main(){
	x := struct{
		mapa	 map[string]int
		slice 	 []int
	}{
		mapa: map[string]int{
			"Tarcisio":32,
			"Barbara": 29,
			"Bryan": 	5,
		},
		slice: []int{3400, 2500, 0},
	}
	fmt.Println(x)
}