/*
	- Slices são feitos de arrays
	- Elas são dinamicas, podem mudar de tamanho
	- Sempre que a mudança de tamanho ocorre, um novo array é criado e os dados são copiados
	- É conveniente, mais tem um custo computacional
	- Para otimizar as coisas, podemos utilizar make
	- Estrutura
		make([]T, len, cap)
	- cria-se um array com uma capacidade maxima prédeterminada, é inserido uma quantidade de valores iniciais já tendo em mente um valor maximo de elementos
		- len se refere ao tamanho inicial do slice
		- cap se refere a ao tamanho que o array que armazena o slice terá, quando o array chega ao limite nesse caso um novo array é criado com o dobro do tamanho
	- make é utilizado para resolver em parte o problema de performace do slice 
*/
package main

import "fmt"

func main(){
	/*Primeiro parametro o tipo, segundo o tamanho do slice terceiro o tamanho do array*/
	slice:=make([]int,5,10)
	/*inserção dos 5 valores*/
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	fmt.Println(slice, len(slice), cap(slice)) 

	slice = append(slice, 6)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 7)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 8)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 9)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 9)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 10)
	fmt.Println(slice, len(slice), cap(slice))
}