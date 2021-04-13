/*
 Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
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

	for x:= 0; x < NumberGoROutines; x++{
		go func(){
			v := contadorRaceCondition
			/*mudança de contexto*/
			runtime.Gosched()
			v++
			contadorRaceCondition = v
			waitgroup.Done()
		}()
		fmt.Println("Goroutines inc:", runtime.NumGoroutine())
	} 
	waitgroup.Wait()
	fmt.Println("Goroutines ativas: ", runtime.NumGoroutine())
	fmt.Println("Valor fim contador: ", contadorRaceCondition)
}