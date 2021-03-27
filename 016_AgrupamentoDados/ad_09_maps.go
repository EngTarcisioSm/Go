/*
	- Maps são uma estrutura de dados do tipo chave valor (key:value)
	- Maps não são ordenados
	- Possui uma performace de busca muito boa
*/
package main

import "fmt"

func main(){

	/*criação de um map com a inserção de alguns valores*/
	amigos := map[string]int{
		"Barbara": 12345678,
		"Bryan": 87654321,
		"Tarcisio": 11223344,
	}

	fmt.Println(amigos)

	/*adicionando um item ao map*/
	amigos["Janete"] = 55667788
	fmt.Println(amigos)

	/*acesso a item específico de map*/
	fmt.Println(amigos["Bryan"])

	/*
		Ocorre uma situação que quando um item de map é acessado e ele não existe é retornado zero. Entretanto uma dúvida fica, returnou zero pois não existe ou pq o valor daquela chave realmente é zero? Para isso no ato da consulta é possivel inserir uma segunda variavel, a linguagem go nessa segunda varivel retorna um booleano, true para o item existir false para o item não existir 
	*/
	nome, tem := amigos["Juca"]
	nome2, tem2 := amigos["Tarcisio"]
	fmt.Println(nome, tem)
	fmt.Println(nome2, tem2)

	/*uma forma de se utilizar o segundo parametro */
	if nome3, tem3 := amigos["Barbara"]; !tem3 {
		fmt.Println("não tem")
	} else {
		fmt.Println(nome3)
	}
}