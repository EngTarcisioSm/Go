/*
	existe a opção default quando nenhuma das opções é atendida
*/
package main

import "fmt"

func main(){
	numero:= 5

	switch numero{
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println(4)
		default:
			fmt.Println("Nenhuma das opções acima")
	}
}