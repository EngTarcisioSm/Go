/*
- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
- Tire estes números do canal e demonstre-os.
*/
package main

import (
	"fmt"
	"sync"
)



func main(){
	c := runGoRoutines()
	go func(){
		for x := range c{
			fmt.Println(x)
		}
	}()
}

func runGoRoutines()chan int{
	canal := make(chan int)
	var wg sync.WaitGroup

	for i:=0; i<10;i++{
		wg.Add(1)
		go func(x int, a chan int){
			a <- x
		}(i,canal)
		wg.Done()
	}
	wg.Wait()
	return canal
}

