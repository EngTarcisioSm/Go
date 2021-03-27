/*
	Põe na tela: O unicode point de todas as letras maiusculas do alfabeto, três vezes cada
*/
package main

import "fmt"

func main(){
	for i := 0x41; i <= 0x5A; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tU+00%d %c\n",i,i)
		}
	}
}