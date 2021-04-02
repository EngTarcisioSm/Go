/*
	- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
	- Passe um valor do tipo slice of int como argumento para a função;
	- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
	- Passe um valor do tipo slice of int como argumento para a função.
*/

package main

import "fmt"

func main(){
	nums := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(soma1(1,2,3,4,5,6,7,8))
	fmt.Println(soma2([]int{1,2,3,4,5,6,7,8}))

	fmt.Println(soma1(nums...))
	fmt.Println(soma2(nums))
}

func soma1 (x...int) int{
	var total int
	for _, a := range x{
		total += a
	}
	return total
}

func soma2(x []int) int{
	var total int
	for _, a := range x{
		total += a
	}
	return total
}
