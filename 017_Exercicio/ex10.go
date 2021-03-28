/*
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
	- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
	- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
	- Demonstre todos esses valores e seus índices.
*/
package main

import "fmt"

func main(){
	person:= map[string][]string{
		"Melo_Tarcisio":[]string{"programar", "jogar"},
		"Melo_Bryan":[]string{"Assistir besteira","comer kinder ovo"},
		"Vasconcelos_Barbara":[]string{"Dormir", "Dormir de dia", "Dormir de noite"},
	}

	fmt.Println(person)

	person["Melo_Alipio"] = []string{"Fazer coisas de casa", "Não ficar parado", "Dançar"}
	delete(person,"Melo_Tarcisio")

	for key, value := range person{
		fmt.Println(key)
		for _, value2 := range value{
			fmt.Printf("\t%s\n", value2)
		}
	}
}