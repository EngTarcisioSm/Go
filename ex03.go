/*
	faça um programa que demonstro o funcionamento de if else if e else
*/
package main

import "fmt"

func main(){
	x:= "abacate"

	if x == "melao" {
		fmt.Print("melao")	
	} else if x == "jaca" {
		fmt.Print("jaca")	
	} else {
		fmt.Print("nenhuma")
	}
}