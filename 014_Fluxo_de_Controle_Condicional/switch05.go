/*
	um switch pode ter mais de um valor para analise
*/
package main

import "fmt"

func main(){
	for i := 1; i <= 6; i++ {
		switch i{
			case 1, 2:
				fmt.Println("Opção A")
			case 3, 4:
				fmt.Println("Opção B")
			case 5,6:
				fmt.Println("Opção C")
			default:
				fmt.Println("Opção D")
		}
	}
}