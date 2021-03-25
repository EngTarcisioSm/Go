/*
	dentro do switch é possivel testar uma variavel
*/
package main

import "fmt"

func main(){
	numero:= 3

	switch numero{
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println(4)
	}
}