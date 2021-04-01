/*
	É uma função que terá uso local apenas uma vez e não possui nome
	- muit utilizado em go routine
*/
package main

import "fmt"

func main(){

	x:= 3

	/*a função é criada,gerada sua logica e au final ja é chamadoao final de sua implementação nesse caso o valor x (x)
	
	*/
	func(x int){
		fmt.Println(4*x)
		fmt.Println("ok")
	}(x)
}