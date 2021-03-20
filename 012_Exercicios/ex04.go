/*
	Crie um programa que:
		- Atribua um valor int a uma variável                                                    - ok
		- Demonstr seu valor em decimal, binário e hexadecinal                                   - ok
		- Desloque os bits dessa varial 1 para esquerda, atribua o resultado a outra variavel    - ok
		- Demonstre esta outra variavel em decimal, binário e hexadecimal						 - ok
*/

/*
	A rotação de um numero a esquerda é a mesma coisa de multiplicar um numero por multiplos de base 2 
*/
package main

import "fmt"

var i uint8

func main(){
	i:= 10

	fmt.Printf("%d - %b - %#x\n",i, i, i)

	j := i << 1

	fmt.Printf("%d - %b - %#x\n", j, j, j)

	fmt.Printf("%T - %T\n", i, j)
}