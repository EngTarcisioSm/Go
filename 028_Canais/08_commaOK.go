/*
	-comma ok
	- Em uma situação em que um canal possui um valor apenas (podendo a teoria ser expandioda para mais valores) quando o canal é retirado seu valor um valor zero fica no canal, existe um problema nisso, como saber se esse valor zero, simplesmente é um valor colocado ou um valor padrão de um canal sem nenhum valor. para isso existe o comma ok.
	Toda vez que um valor for inserido e em outro ponto do código for retirado o ok retornará true indicando que o valor foi inserido, caso retorne falso isso indica que o valor coletado não foi inserido dentro do canal, podendo ser desconsiderado
*/

/*
	- canais se comportam como queues(filas), fifo (first input, first output)
*/
package main

import "fmt"

func main(){
 	canal := make(chan int)
	 go func(){
		canal <- 42
		close(canal)
	 }()

	 v, ok := <-canal
	/*como o valor foi inserido ok retornará true*/
	 fmt.Println(v, ok)
	
	 v, ok = <-canal
	/*o valor que existia foi retirado na linha anterior, o valor retirado agora não foi inserido logo ok retornará false*/
	 fmt.Println(v, ok)
}