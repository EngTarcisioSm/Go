/*
	- Utilizando como base o exercicio anterior, utilize slicing para demonstrar os valores:
		- Do primeiro ao terceiro item do sice(incluindo o terceiro item!)
		- Do Quinto ao ultimo item do slice (incluindo o ultimo item!)

		- Do terceiro ao penultimo item do slice (incluindo o penultimo item!)
		- Desafio: obtenha o mesmo resultado acuma utilizando a função len() para determinar o penultimo item
*/
package main

import "fmt"

func main(){
	slice := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(slice[:3]) 
	fmt.Println(slice[5:])
	fmt.Println(slice[3:8])
	fmt.Println(slice[3:len(slice)-2])
}