/*
	- Alem da goroutine principal, crie duas outras goroutines.
	- Cada goroutine adicional devem fazer um print separado.
	- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){

	wg.Add(2)

	go print1()
	go print2()

	wg.Wait()
}

func print1(){
	fmt.Println("Bryan")
	wg.Done();
}
func print2(){
	fmt.Println("Tarcísio")
	wg.Done();
}