/*
	- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
	- Demonstre os valores do map utilizando range.
	- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
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

	databank := map[string]pessoa{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	for key, value := range databank{
		fmt.Println(key)
		for _, value2 := range value.sorvete{
			fmt.Printf("\t%s\n",value2)
		}
	}
}