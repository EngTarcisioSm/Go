/*
	Loop for
		- Inicialização, condição e pós
		- Não existe while na linguagem go 
		- goyexample.com
*/
package main

import "fmt"

func main(){

	/*inicial condição incremento*/
	for x:= 0; x < 10; x++{
		fmt.Println(x)
	}
}