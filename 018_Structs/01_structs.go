/*
	- Struct
		- Tipo de dado
		- Vem da palavra inglesa estrutura(em português)
		- Permite armazenar dados com tipos diferentes
*/
package main

import "fmt"

/*
definindo uma estrutura
type nome_estrutura struct{
	nomeVariavel1 tipoDado1,
	nomeVariavel2 tipoDado2,
	.
	.
	.
}
*/
type cliente struct{
	nome      string
	sobrenome string
	fumante   bool
}


func main(){
	/*formas de definir uma struct e inserir valores*/
	/*Forma 1: mais clara de definir "mais recomendada*/
	cliente1 := cliente{
		nome:     "João",
		sobrenome:"Da Silva",
		fumante:  false,
	}
	/*Forma 2*/
	cliente2 := cliente{"Joana", "Pereira",true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}