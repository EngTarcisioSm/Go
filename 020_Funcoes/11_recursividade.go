/*
	- Recursividade: inseption de função, a função chamando a si própria

	- loops podem substituir recursividade

*/

package main

import "fmt"


func main(){
	fmt.Println(fatorial(20))
}

func fatorial(x int) int{
	if x == 1{
		return x
	}
	return x * fatorial(x-1)
}