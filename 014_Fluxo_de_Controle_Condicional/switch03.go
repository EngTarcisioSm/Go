/*
	através de uma palavra reservada fallthrough é possivel fazer com que caso um dos cases do switch seja executado, ele executará o seguinte também
*/
package main

import "fmt"

func main(){
	numero:= 2

	switch numero{
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
			fallthrough
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println(4)
	}
}