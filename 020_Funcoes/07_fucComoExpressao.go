/*
	Atribuir a uma variavel uma função
*/
package main

import "fmt"


func main(){

		x := func(x int){
			fmt.Println(x, "...")
			fmt.Println(x,"->")
		}

		x(30)

}