/*
	- Convergencia
		- Padrão de comunicação, uma forma de se fazer as coisas;
		- Pegar dados de vários pontos e convergir eles para apenas alguns poucos ou um unico

*/
/*
	- Exemplo:
	- Canais par, ímpar, e converge.
        - Func send manda pares pra um, ímpares pro outro, depois fecha.
        - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int) 

	go envia(par, impar)

	go recebe(par, impar, converge)

	for v:= range converge{
		fmt.Println(v)
	}

}

func envia(p, i chan int) {
	x := 100
	for n := 1; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int){
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		for v:= range p{
			c <- v
		}
		wg.Done()
	}()

	wg.Add(1)
	go func(){
		for v:= range i{
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)

}
