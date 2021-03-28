/*
- Crie uma slice usando make que possa conter todos os estados do Brasil.
    - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice *sem utilizar range.*
*/
package main

import "fmt"

func main(){

	states:= make([]string, 26, 50)
	states[0],states[1],states[2],states[3],states[4],states[5],states[6],states[7],states[8],states[9],states[10],states[11],states[12],states[13],states[14],states[15],states[16],states[17],states[18],states[19],states[20],states[21],states[22],states[23],states[24],states[25] = "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
	fmt.Println(len(states), cap(states))

	for x:=0;x<26;x++{
		fmt.Println(states[x])
	}
}