/*
	- Utilize atomic para consertar a condição de corrida do exercício #3.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


func main(){
	fmt.Println(".......................")
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println(".......................")

	var contadorRaceCondition int64 
	contadorRaceCondition = 0

	NumberGoROutines := 1000;

	var waitgroup sync.WaitGroup
	waitgroup.Add(NumberGoROutines)

	for x:= 0; x < NumberGoROutines; x++{
		go func(){
			atomic.AddInt64(&contadorRaceCondition, 1)
			/*mudança de contexto*/
			runtime.Gosched()
			waitgroup.Done()
		}()
		fmt.Println("Goroutines inc:", runtime.NumGoroutine())
	} 
	waitgroup.Wait()
	fmt.Println("Goroutines ativas: ", runtime.NumGoroutine())
	fmt.Println("Valor fim contador: ", contadorRaceCondition)
}