
package main

import "fmt"

func main(){
	/*string*/
	s := "Tarcísio Souza de Melo"
	/*raw*/
	sRaw := `String                     
	bagunçada            	!`
	fmt.Println(s)
	fmt.Println(sRaw)

	/*byte slice*/
	sSlice := []byte(s)

	/*loop is not yet explained*/
	for _, v:= range sSlice{
		/*
			%v  - returns the value in the decimal variable
			%T  - Returns the type of the variable
			%#V - Returns the UTF-8 value
			%#x - returns the hex value
		*/
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}
	v:= "Tarcisio 戦"
	fmt.Println()
	/*Capturing UTF-8 characters*/
	/*len() retorna o tamanho de uma string e o tamanho de outras estruturas*/
	for i:= 0; i < len(v); i++ {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v[i],v[i],v[i],v[i],v[i])
	}

}
