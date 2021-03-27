/*
	Quando se utiliza um append() para copiar um trecho de slice para outro slice, não ocorre a cópia do array que forma o primeiro slice para o segundo. O mesmo slice é copiado, isso não é alguma feacture mais sim um problema.
	Tendo ciencia desse problema caso seja necessário a cópia se faz necessário copiar item por item ou utilizar o proprio array primário
	É importante mensionar que mesmo que o segundo slice que tem trecho do primeiro, o tamanho do segundo e de um tamanho x só que a capacidade dele é maior, os demais numeros do array compartilhado fica escondido
*/
package main

import "fmt"

func main(){
	slice_1 := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(slice_1)
	slice_2 := append(slice_1[:2],slice_1[5:]...)

	/*ocorreu um fatiamento no slice 1 para o slice 2 utilizar o mesmo array do slice*/
	fmt.Println(slice_1)
	fmt.Println(slice_2)

	/*É possivel verificar que a modificação de um valor no array 1 modifica o array 2*/
	slice_1[0] = 100
	fmt.Println()
	fmt.Println(slice_1)
	fmt.Println(slice_2)

}