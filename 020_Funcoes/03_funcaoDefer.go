/*
	defer
		- funcionalidade que faz com que uma linha de código seja executada após todos os processos terem se encerrado
		- caso exista mais de uma linha com defer ocorrerá que funcionará como uma pilha ou seja, o primeiro a entrar será o ultimo a sair 

		- torna o código mais modular
		- onde encerrar o codeblock "}" é o momento em que o defer será executado, caso tenha um retorno antes do codeblock ele será executado uma linha antes do code block 

		-defer em portuguÊs significa adiar 
*/
package main

import "fmt"

func main(){

	defer fmt.Println("Primeiro")
	fmt.Println("Segundo")

}