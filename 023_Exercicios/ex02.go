/*
	- Ao utilizar ponteiros de struct o correto seria utilizar (*nomestruct).conteudo entretanto existe uma exceção para o uso podendo ser abreviado em nomestruct.conteudo
*/
package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade int
}

func main(){
	tarcisio := pessoa{
		nome: "Tarcisio",
		sobrenome: "Souza de Melo",
		idade: 32,
	}

	mudeMe(&tarcisio)

	fmt.Println(tarcisio)

}

func mudeMe(x *pessoa){
	x.idade++
}