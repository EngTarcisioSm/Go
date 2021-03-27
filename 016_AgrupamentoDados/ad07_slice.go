/*
	Slice multidimensional
		- Semelhante a uma matriz ou uma planilha no excel
*/
package main

import "fmt"


func main(){
	/*Slice multidimensional*/
	ss := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	fmt.Println(ss)

	/*POssivel pegar trechos*/
	/*linha*/
	fmt.Println(ss[2])
	/*item*/
	fmt.Println(ss[1][2])

	/*os slices não necessitam de ter o mesmo tamanho*/
}