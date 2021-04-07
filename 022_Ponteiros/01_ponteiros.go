/*
	- Todo valor armazenado no computador ele é alocado em uma memoria, e essa memória possui um endereço, tendo o endereço dessa memoria é possivel acessar o valor contido na memoria 
	- Segue a mesma Lógica da linguagem C
*/

package main


import "fmt"

func main(){
	x := 10
	y := &x /*passando o endereço da variavel x*/

	fmt.Println(*y) /*indicando o valor presente no endereço de memoria que y guarda*/

	fmt.Printf("%T\n",y) /*indicando o tipo de int*/

	fmt.Println(y, &x) /*indicando o endereço de x e o valor armazenado em y*/

	*y = 12

	fmt.Println(x) /*ao alterar o conteudo que y aponta o valor de x é alterado também */
}