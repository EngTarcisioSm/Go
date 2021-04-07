/*
	-Ponteiros servem para 
		- melhorar performace, toda vez que uma variavel é fornecida como atributo ela é passada uma cópia e essa cópia possui um custo computacional 
*/

package main

import "fmt"

func main(){
	x :=0

	recebeponteiro((&x))
}

/*função que recebe um ponteiro como parametro */

func recebeponteiro(x *int){
	*x++
	fmt.Println("Na função: ", *x)
}