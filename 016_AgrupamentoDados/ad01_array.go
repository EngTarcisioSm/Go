/*
	Maneiras diferente que se tem para se agrupar informações dentro do programa 
	Quatro tipos principais
		- arrays: menos utilizado no dia a dia 
			- possui apenas um tipo
		- slices: mais utilizado no dia a dia entretanto este é construido sobre arrays
		- maps 
		- structs
*/
package main

import "fmt"

/*definição de um array*/
var x[10]int
var y[5]int

func main(){

	/*definição curta*/
	array:= [5]int{1,2,3,4,5}
	//inserção de um valor sobre o array
	x[0] = 12
	x[1] = 100
	x[2] = 2

	//impresão
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Println(array)
	
	//arrays com o mesmo tipo de dado se não forem do mesmo tamanho, não serão do mesmo tipo
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	//ver o tamanho de um array
	fmt.Println(len(x))

	//array é utilizado como desejamos utilizr par atimizar extremamente a memoria  
}