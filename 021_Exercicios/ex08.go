/*
	- Crie uma função que retorna uma função.
	- Atribua a função retornada a uma variável.
	- Chame a função retornada.
*/
package main

import "fmt"


func main(){
	x := funcao1()
	fmt.Println(x(10))

	fmt.Println(funcao1()(1))
}

func funcao1() func(int)int{
	return func(x int) int{
		return x+1
	}
}