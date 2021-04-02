/*
	- Atribua uma função a uma variável.
	- Chame essa função.
*/
package main

import "fmt"


func main(){
	x := func(s string){
		fmt.Println(s)
	}

	x("Tarcísio Souza de Melo")
	x("Bryan William Vasconcelos de Melo")
	fmt.Printf("%T\n",x)
}