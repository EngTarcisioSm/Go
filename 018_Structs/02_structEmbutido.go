/*
	Struct embutido nada mais é do que um struct dentro de outro struct
*/
package main

import ("fmt"
)
type pessoa struct{
	nome	string
	idade	int
}

type profissao struct{
	pessoa
	titulo	string
	salario	int
}

func main(){
	pessoa1 := profissao{
		pessoa: pessoa{
			nome: "Alfredo",
			idade: 38,
		},
		titulo:"pizzaiolo",
		salario: 10000,
	}
	/*outra forma de se definir um valor */
	pessoa2 := profissao{pessoa {"Vanderlei", 50}, "Político", 10000000}

	fmt.Println(pessoa1)

	/*Acessando apenas um campo da struct*/
	fmt.Println(pessoa1.titulo)

	fmt.Println(pessoa1.pessoa.idade)

	/*quando existe uma estrutura dentro de outra e o nome de um membro da estrutura mais interna não se repete nas mais externas é possivel acessar diretamente*/
	fmt.Println(pessoa1.nome)

	fmt.Println(pessoa2)

	
}