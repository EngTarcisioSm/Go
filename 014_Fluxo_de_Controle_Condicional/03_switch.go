
package main

import "fmt"

func main(){
	x := 10

	/*Executa o primeiro caso com resultado true*/
	switch{
		case x < 5:
			fmt.Println("1")
		case x > 15:
			fmt.Println("2")
		case x > 5:
			fmt.Println("3")

	}

}