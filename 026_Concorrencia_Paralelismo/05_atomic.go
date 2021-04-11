/*
	- sync/atomic
	- sem race condition (condição de corrida)
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main(){
	fmt.Println("CPU: ", runtime.NumCPU())

	var contador int64
	contador = 0
	totaldegoroutines := 15

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i:=0; i<totaldegoroutines; i++{
		go func(){
			atomic.AddInt64(&contador,1)
			runtime.Gosched() //muda de contexto
			fmt.Println("contador: \t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("valor final: ", contador)
}