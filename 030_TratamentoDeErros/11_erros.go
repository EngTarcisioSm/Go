/*
	- variação um pouco diferente do problema anterior
	- informações extras:
		- http://golang.org/src/pkg/bufio/bufio.go			
		- http://golang.org/src/pkg/io/io.go
*/
package main

import (
	"errors"
	"fmt"
	"log"
)

//ajuda no caso de uma função que possui varias funcionalidades, podendo ter diversos pontos de retorno de erro na função permitindo um código mais limpo 
var ErrNorgatheMath = errors.New("norgate math: square root of negative number")

func main(){
	//Demonstra o tipo da variavel abaixo 
	fmt.Printf("%T\n", ErrNorgatheMath)

	_, err := sqrt(-10)
	if err != nil{
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64,error){
	if f < 0 {
		return 0, ErrNorgatheMath
	}
	return 42, nil
}