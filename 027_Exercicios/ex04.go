/*
	- Utilize mutex para consertar a condição de corrida do exercício anterior.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main(){
	fmt.Println(".......................")
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println(".......................")

	contadorRaceCondition := 0

	NumberGoROutines := 1000;

	var waitgroup sync.WaitGroup
	waitgroup.Add(NumberGoROutines)

	var mu sync.Mutex

	for x:= 0; x < NumberGoROutines; x++{
		go func(){
			mu.Lock()
				v := contadorRaceCondition
				/*mudança de contexto*/
				runtime.Gosched()
				v++
				contadorRaceCondition = v
			mu.Unlock()
			waitgroup.Done()
		}()
		fmt.Println("Goroutines inc:", runtime.NumGoroutine())
	} 
	waitgroup.Wait()
	fmt.Println("Goroutines ativas: ", runtime.NumGoroutine())
	fmt.Println("Valor fim contador: ", contadorRaceCondition)
}