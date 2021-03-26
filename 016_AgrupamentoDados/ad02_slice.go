/*
	- Slice é um tipo de dado composto utilizando os dados primitivos da linguaagem
	- diferente do array que se define um valor de tamanho em sua inicialização o slice que possui definição de tamanho inicial
	- slice é uma estrutura que pode crescer indefinidamente conforme a necessidade do programa
	- sua inicialização é semelhante a do array com a diferença que não se define um tamanho maximo
*/
package main

import (
	"fmt"
)

/*definição de um slice*/

var slice1[]int

func main(){
	/*definição de array  com :=*/
	slice2 := []int{1,2,3,4,5}

	fmt.Printf("%T, %d\n",slice1,slice1)
	fmt.Printf("%T, %d\n",slice2,slice2)

	fmt.Println(slice2[0])
	slice2[0] = 1000
	fmt.Println(slice2[0])

	/*quando se vai inserir um novo elemento em um slice o que ocorre é que ele copia o slice antigo para um novo e adiciona o valor nesse novo */
	slice3 := append(slice2,555)
	fmt.Println(slice3[5])
	/*o slice que irá receber o novo slice com a quantidade de valores maior pode ter o mesmo nome*/
	slice3 = append(slice3, 101)
	fmt.Println(slice3[6])
}