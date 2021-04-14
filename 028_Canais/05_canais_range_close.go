/*
	- usar range em canais é comum 
	- quando se usa um canal e esvazia ele, o mesmo aguarda novos valores, se o programa encerrar com um canal aberto aguardando dados gera-se um erro, para isso é recomendavel o encerramento dos canais 
 */

package main

import "fmt"

func main(){
	c:= make(chan int)

	go meuloop(10, c)

	for v := range c{
		fmt.Println(v)
	}
	close(c)
}

func meuloop (t int, s chan<- int){
	for i:= 0; i < t; i++{
		s <- i
	}
}