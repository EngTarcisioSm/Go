/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/
package main

import (
	"fmt"
)

type pessoa struct{
	nome		string
	sobrenome	string
	idade		int
}

func dados(Person pessoa) {
	fmt.Println(Person.nome, Person.sobrenome, Person.idade)
}


func main(){

	pessoaBarbara := pessoa{
		nome: "Barbara",
		sobrenome: "Vasconcelos",
		idade: 29,
	}
	dados(pessoaBarbara)

}