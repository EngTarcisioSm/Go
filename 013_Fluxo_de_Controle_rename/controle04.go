/*
	na linguagem go não existe o loop while entretanto é possivel modificando o for fazer ter o mesmo comportamento de um while
*/
package main

import "fmt"

func main(){
	
	x:= 0 //inicialização fora do loop

	for x < 10{ //condição de teste
		fmt.Printf("x:%d\n",x)
		x++ //incremento da variavel
	}
}