/*
	Desenrolando um slice para passar ele como argumento para uma função
*/
package main

import "fmt"

func main(){
	valores := []int{1,2,3,4,5,6,7,8}
	/*
		não é possivel passar diretamente o valor de um slice para uma função como apresentado é necessário utilizar o operador ...
	*/
	fmt.Println(soma(valores...))
}

/*quando uma função utiliza ... para receber parametros ela pode receber infinitos valores como também valor nenhum, neste ultimo caso retornando o valor zero*/
/*se uma função utiliza o operador ... ele tem de ser o ultimo parametro caso tenha mais de um parametro*/
func soma(x...int) int{
	soma := 0
	for _, a := range x{
		soma += a
	}
	return soma
}