/*
	- Select é como switch, só que pra canais, e não é sequencial.
	- "A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
		- o select bloqueia até receber algum dado que bate com os valores esperado por ele 
		- caso se tenha mais de um caso valido será escolido um aleatorio não podendo se ter controle sobre o mesmo ​
	- Na prática:
		- Exemplo 1:
			- Duas go funcs enviando X/2 numeros cada uma pra um canal
			- For loop X valores, select case ←x
		- Exemplo 2:
			- Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
			- Func 2 for infinito, select: case envia pra canal, case recebe de quit
		- Exemplo 3:
			- Chans par, ímpar, quit
			- Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
			- Func receive é um select entre os três canais, encerra no quit
			- Problema!
*/

/*
	- Exemplo:
		- Duas go funcs enviando X numeros cada para um canal 
		- utilizando o select
			for loop x values, select case <- x
*/
package main

import "fmt"

func main(){
	x:= 500
	a := make(chan int)
	b := make(chan int)

	go func(x int){
		for i:=0; i< x; i++{
			a <- i
		}
	}(x/2)
	
	go func(x int){
		for i:=0; i< x; i++{
			b <- i
		}
	}(x/2)

	for i:=0; i < x; i++{
		select{
			case v := <-a:
				fmt.Println("canal 1: ", v)
			case v := <-b:
				fmt.Println("canal 2: ", v)
		
		}
	}

}