/*
	switch  baseado no tipo da variavel
*/
package main

import "fmt"

/**/
var numero interface{}

func main(){
	numero = 3

	switch numero.(type){
		case int:
			fmt.Println(1)
		case bool:
			fmt.Println(2)
		case string:
			fmt.Println(3)
		case float64:
			fmt.Println(4)
	}
}