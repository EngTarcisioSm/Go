/*
	Utilização de continue dentro de for
		- Dado uma condição se esta tiver como processo um continue a execução naquela iteração é pulada não executando o loop até o final e pulando para a proxima iteração
*/
package main

import "fmt"

func main(){
	for i:= 0; i <= 100; i++{
		if i % 2 != 0{
			continue
		}
		fmt.Printf("%d\n", i)
	}
}