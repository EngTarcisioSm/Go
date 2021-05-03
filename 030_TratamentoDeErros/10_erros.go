/*
	- errors.New("norgate math: sqare root of negative number")
		- cria a mensagem de erro
	- log.Fatalln(err)
		- faz a mensagem retornar como um log de erro do nada e hora 

*/
package main

import(
	"errors"
	"log"
)

func main(){
	_, err := sqrt(-10)
	if err != nil{
		/*
			caso ocorra um erro será utilizado um Fatalln faz o programa retornar exit status 1
		*/
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error){
	if f < 0 {
		return 0, errors.New("norgate math: sqare root of negative number")
	}
	return 42, nil
}