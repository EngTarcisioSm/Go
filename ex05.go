/*
	switch com switch statement com nome de esporte favorito
*/
package main

import "fmt"

func main(){
	esporteFavorito := "futebol"

	switch esporteFavorito{
		case "volei":
			fmt.Print("Volei")
		case "Formula1":
			fmt.Print("corrida")
		case "Basquete":
			fmt.Print("Basquete")
		default:
			fmt.Print("nenhum dos anteriores")
	}
}