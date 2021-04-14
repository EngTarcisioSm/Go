/*
- Exemplo 3:
			- Chans par, ímpar, quit
			- Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
			- Func receive é um select entre os três canais, encerra no quit
			- Problema!
				- ao realizar a ultima operação é retornado um valor falso 0 impar, problema tratado na proxima aula
*/
package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumeros(par, impar, quit)
	receive(par, impar, quit)

}

func mandaNumeros(par, impar chan int, quit chan bool) {
	total := 500

	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("Numero par: ", v)
		case v := <-impar:
			fmt.Println("Numero impar: ", v)
		case <-quit:
			return
		}
	}

}
