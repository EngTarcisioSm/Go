/*
	- Usando uma literal composta:
		- Atribua 10 valores a ela
		- Utilize range para demonstrar todos estes valores.
		- Utilize format printing para demonstrar seu tipo.
*/
package main

import "fmt"

func main(){

	slice := []int{1,2,3,4,5,6,7,8,9,10}

	for _, value := range slice{
		fmt.Println(value)
	}
	fmt.Printf("%T",slice)
}