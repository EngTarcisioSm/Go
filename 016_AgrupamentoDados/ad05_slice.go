/*
	append()
		- utilizado para inserir dados a um slice
		- ao utilizar o append em um slice o primeiro parametro é o slice que deseja que seja adicionado (por mais que não ocorra a inserção mais sim uam cópia), o segundo parametro que podem ser varios, é o valor que será adicionado ao slice, só pode ser adicionado a um slice conteudo do mesmo tipo dos que foram inseridos anteriormente
		- quando se deseja que um slice seja inserido em outro deve-se utilizar ...(chamado de unfurl e na documentação enumeration) que meio que quebra o slice em suas partes unitárias permitindo assim a inclusão
*/
package main

import (
	"fmt"
)

func main(){
	slice1 := []int{1,2,3,4,5}
	fmt.Println(slice1);
	/*aidionando varios itens ao slice*/
	slice1 = append(slice1, 6,7,8,9)
	fmt.Println(slice1);
	slice2 := []int{10,11,12,13,14,15,16}
	/*adicionando um slice (slice2) a outro slice (slice1)*/
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}