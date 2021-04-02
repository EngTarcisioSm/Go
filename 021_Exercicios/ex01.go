/*
	- Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
*/

package main

import "fmt"

func main(){

	fmt.Println(retint())
	fmt.Println(retIntStr())
}

func retint () int {
	return 32
}

func retIntStr() (int, string){
	return 32, "Tarcísio"
}