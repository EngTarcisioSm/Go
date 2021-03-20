package main

import "fmt"

func main() {
	//interpreted string literal
	a := "Tarcisio \t Souza \t de Melo\n" 
	// raw string literal
	b := `Bryan              William            Vasconcelos de Melo` 
	
	fmt.Printf(a)
	fmt.Printf(b)
	fmt.Println()
}