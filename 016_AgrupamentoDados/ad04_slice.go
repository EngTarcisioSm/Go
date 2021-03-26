/*
	Fatiando um slice
*/
package main

import "fmt"

func main(){

	sabores:=[]string{"pepperoni", "mussarela", "abacaxi", "4 queijos", "marguerita"}
	/*
		Quando se efetua um  fatiamento de um slice, o primeiro valor da um intervalo fechado ou seja ele será inclusõ na resposta final, enquanto o segundo valor é o intervalor aberto, ou seja ele não estará presente na resposta final
	*/
	fatia := sabores[0:2]

	fmt.Println(fatia)
	/*
		- Formas de inserir valores para se efetuar o slice 
		[;x] -> vai da posição zero até a posição x, com a posição x não inclusa na resposta
		[x;] -> vai da posição x que esta incluso até o final do slice (pegará todos os valores a partir do indice x)
		[y;len(slice)] -> pegará os itens da posição y até o final do slice pois len retornará o valor do tamanho do slice o qual é tamanhoMax+1 
	*/

	/*acessando todos itens de um slice sem usar range */

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

	/*
		Deletando o item de um slice
			- utiliza-se o mesmo metodo de se adicionar um item 
			- utiliza-se o simbolo ... no segundo parametro de append
	*/

	/*remover o item 2 do slice abaixo*/
	sabores = append(sabores[:1],sabores[3:]...)
	fmt.Println(sabores)
}