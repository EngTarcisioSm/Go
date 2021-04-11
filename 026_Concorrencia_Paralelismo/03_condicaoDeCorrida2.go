/*
	- Varias go funcs acessaram uam variavel ao mesmo tempo acarretando o não cumprimento de esperar aquela variavel ser incrementada para ser utilizada, ocorreu o acesso em memória ao mesmo tempo, não respeitando a logica desejada
		func main(){
		//Apenas para teste 

		contador := 0
		totaldegoroutines := 10
		var wg sync.WaitGroup
		wg.Add(totaldegoroutines)

		for i := 0; i < totaldegoroutines; i++{
			go func(){
				v := contador
				//yield - permite ao sistema ir tratar outra função
				runtime.Gosched()
				v++
				contador = v
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Println(contador)
		}

		- é possivel verificar a existencia de condições de corrida no codigo ao executa-lo digitar go run -race nomeArquivo.go, caso seja encontrado uma condicao de corrida o código pode aprensentar comportamento anomalo  
		
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

	for i := 0; i < totaldegoroutines; i++{
		go func(){
			v := contador
			//yield - permite ao sistema ir tratar outra função
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines inc: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines end: ", runtime.NumGoroutine())
	fmt.Println("Valor fim: ",contador)

}