/*
	-Mutex:
		- faz com com um recurso (variavel, ...) só possa ser utilizado por outras goroutines quando a primeira que solicitou o uso tenha finalizado 
	- é necessário criar uma estrutura com o tipo desejado 
	- é necessário uma função que bloquei o acesso a essa estrura e outro que libere o acesso a essa estrutura 

	- ao verificar a condição de corriga do go run -race, não é comunicado nenhuma informação indicando que não ha possibilidade de comportamento erratico do sistema
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main(){
	//Apenas para teste 
	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 1000
	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	var mu sync.Mutex //criação da mutex

	for i := 0; i < totaldegoroutines; i++{
		go func(){
			mu.Lock() //travamento do recurso ou processos abaixo
				v := contador
				//yield - permite ao sistema ir tratar outra função
				runtime.Gosched()
				v++
				contador = v
			mu.Unlock() //liberação dos recursos 
			wg.Done()
		}()
		fmt.Println("Goroutines inc: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines end: ", runtime.NumGoroutine())
	fmt.Println("Valor fim: ",contador)
}