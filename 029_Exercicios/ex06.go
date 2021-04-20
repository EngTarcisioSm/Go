/*
	- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
*/
package main

import (
	"fmt"
)

func main(){
	y := make(chan int)
	to100(y)
	printNumbers(y)
	
}

func to100(x chan int){
	go func(a chan int){
		for i:= 0; i <100; i++{
			a <- i
		}
		close(a)
	}(x)
}

func printNumbers(x <-chan int){
	for{
		select{
		case value,ok:= <-x:
			if(ok){
				fmt.Println(value)
			} else {
				return
			}
		}
	}
}