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
		case v,ok := <-par:
			if ok {
				fmt.Println("Numero par: ", v)
			}
		case v, ok := <-impar:
			if ok {
				fmt.Println("Numero impar: ", v)
			}
		case v, ok := <-quit:
			if ok && v {
				fmt.Println("Terminou...")
			} else {	
				fmt.Println("Zebra")
			}
			return
		}
	}

}
