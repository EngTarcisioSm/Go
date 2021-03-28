/*
	- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
	- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/
package main

import "fmt"

type pessoa struct{
	nome		string
	sobrenome	string
	sorvete		[]string
}

func main(){
	pessoa1 := pessoa{"Tarcisio","Souza de Melo",[]string{"Creme","Napolitano","Chocolate"}}
	pessoa2 := pessoa{"Bryan","William Vasconcelos de Melo",[]string{"Chocolate", "Morango"}}

	fmt.Println(pessoa1.nome)
	for _,sabores := range pessoa1.sorvete{
		fmt.Printf("\t%s\n",sabores)
	}

	fmt.Println(pessoa2.nome)
	for _,sabores := range pessoa2.sorvete{
		fmt.Printf("\t%s\n",sabores)
	}
}