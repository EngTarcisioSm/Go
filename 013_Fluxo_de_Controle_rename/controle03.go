/*
	Nested loop - Repetição Hierarquica
		- Repetição dentro de outra repetição
*/
package main

import "fmt"

func main(){
	for hora:= 0; hora <= 12; hora++{
		for minutos:=0; minutos < 60; minutos++{
			for segundos:=0; segundos < 60; segundos++{
				fmt.Printf("%dH:%dM:%dS\n",hora, minutos, segundos)
			}
		}
	}
}