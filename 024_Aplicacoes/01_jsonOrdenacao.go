/*
	- JSON: Java Script Object Notation

	- O pacote "enconding/json" codifica e decodifica JSON

	- marshal significa organizar, ordenar

	- Quando se vai mexer com pacotes em Go tudo que é em minusculo não é visivel externamente, para estar visivel externamente a primeira letra obrigatoriamente deve ser maiusculo de tudo que se deseja externar 

	- Para se trabalhar com JSON em go todos os atributos bem como estruturas devem ter pelo menos a primeira letra maiuscula 
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct{
	Nome			string
	Sobrenome		string
	Idade			int
	Profissao		string
	Contabancaria	float64
}

func main(){

	jamesbond := Pessoa{
		Nome: "James",
		Sobrenome: "Bond",
		Idade:	40,
		Profissao: "espião",
		Contabancaria: 10000000,
	}

	darthvader := Pessoa {
		Nome: "Anaquin",
		Sobrenome: "Skywalker",
		Idade: 60,
		Profissao: "terrorista intergalatico",
		Contabancaria: 50000000000000,
	}

	j, err1 := json.Marshal(jamesbond)
	d, err2 := json.Marshal(darthvader)

	if err1 != nil{
		fmt.Println("Erro James Bond")
	} else {
		fmt.Printf("%T, %s\n", j, j)
	}
	if err2 != nil{
		fmt.Println("Erro Darth Vader")
	} else {
		fmt.Printf("%T, %s\n", d, d)
	}

}