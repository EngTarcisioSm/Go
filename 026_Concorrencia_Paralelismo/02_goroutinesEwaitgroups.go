/*
	- GoRoutines são algo parecido com threads
	- threads:
		forma de fazer um código ser segmentado para que este seja rodado de forma concorrencialmente

	- quando se coloca "go" na frente de um nome de uma função ela passa a ser uma goroutine
		- colocando apenas go na frente da função não resolverá


	-sync.WaitGroup
		- grupo de espera
		- informa ao sistema que existe go routines que devem ser executadas
		- Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
		- como usar
			- func Add: "Quantas goroutines?" (no inicio da main)
			- func Done: "Deu!" (dentro da go routine )
			- func Wait: "Espera todo mundo terminar." ( ao fim da função main)


*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main(){	
	
	// /*rodando de forma sequencial*/
	// func1() //apenas com a função normal 
	// func2()


	/*informa o numero de cores */
	fmt.Println("......................")
	fmt.Println("numero de cores: ", runtime.NumCPU()) //apenas curiosidade
	fmt.Println("......................")
	fmt.Println("numero de goroutines iniciais: ", runtime.NumGoroutine())
	fmt.Println("......................")

	wg.Add(2)

	go func1()
	go func2()


	fmt.Println("......................")
	fmt.Println("numero de goroutines finais: ", runtime.NumGoroutine())
	fmt.Println("......................")
	wg.Wait()
}

// função normal
// func func1(){
// 	for i:=0;i<10;i++{
// 		fmt.Println("func1 :", i)
// 	}
// }

//função preparada para go routine
func func1(){
	for i:=0;i<100;i++{
		fmt.Println("func1 :", i)
	}
	wg.Done()//feito
}

// func func2(){
// 	for i:=0;i<10;i++{
// 		fmt.Println("func2 :", i)
// 	}
// }

func func2(){
	for i:=0;i<100;i++{
		fmt.Println("func2 :", i)
	}
	wg.Done()
}