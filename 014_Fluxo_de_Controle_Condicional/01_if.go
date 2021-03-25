/*
	Declaração if
*/
package main

import "fmt"

func main(){

	x:=10

	if x > 100 {
		fmt.Println("Sim")
	}

	if !(x > 100){
		fmt.Println("Não")
	}

	// É possivel inicializar a variavel que será checada no processo na mesma linha em que se verifica a condição do if
	if y:=112; y != 10{
		fmt.Println(y)
	}
}