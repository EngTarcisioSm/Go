/*
	Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
	- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
package main

import "fmt"

func main(){
	slice := [][]string{
		[]string{"Tarcisio", "Melo", "programar"},
		[]string{"Bryan", "Melo", "assistir besteira"},
		[]string{"Barbara","Vasconcelos", "dormir"},
	}

	for index,person:=range slice{
		fmt.Println(index,person)
	}
}