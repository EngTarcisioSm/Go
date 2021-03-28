/*
	Utilizando uma literal composta
		- Crie um array que suporte 5 valores do tipo int
		- Atribua valores aos seus índices 
		- Utilizando range demonste os valores do array
		- Utilizando format printing, demonstro o tipo do array
*/
package main

import "fmt"

func main(){
	x := [5]int{10,20,30,40,50}

	for _,value := range x{
		fmt.Println(value)
	}
	fmt.Printf("%T",x)
}