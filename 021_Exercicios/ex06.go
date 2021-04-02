/*	
	- Crie e utilize uma função anônima.
*/
package main

import "fmt"


func main(){
	func(s string){
		fmt.Println(s)
	}("Tarcísio Souza de Melo")
}