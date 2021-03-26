/*
função range dentro de slice
*/
package main 

import "fmt"

func main(){
	/*range quer dizer alcance auxilia em percorrer um slice com auxilio do laço for*/
	slice := []string{"banana", "jaca", "pera", "uva", "melancia"}

	/*no uso do range é retornado indice da estrutura slice e o valor da posição com o mesmo indice*/
	for indice, valor := range slice{
		fmt.Printf("Indice %d: %s\n",indice, valor)
	}
}