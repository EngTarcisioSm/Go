/*
	Demonstre o resultado da divisão por 4 de todos os numeros entre 10 a 100
*/
package main

import "fmt"

func main(){
	resto := 0
	for i := 10; i <= 100; i++ {
		resto = i % 4
		fmt.Printf("i -> %d\n",resto)
	}
}