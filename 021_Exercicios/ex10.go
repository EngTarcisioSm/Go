/*
	- Demonstre o funcionamento de um closure.
	- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/
package main

import "fmt"


func main(){
	x := mainScope(20)

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	

}

func mainScope(x int) func() int{
	return func()int{
		x++
		return x 
	}
}