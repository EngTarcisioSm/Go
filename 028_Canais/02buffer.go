/*
	- buffers também bloqueiam
	- possui um espaço para armazenamento

	-em geral não se utiliza buffer chanel (contar com sincronização e não esperar buffers)
*/
package main

import "fmt"

func main(){
	//buffer de tamanho 1
	canal := make(chan int,1)
	canal <- 42
	//caso tente inserir mais de um valor nesse buffer que tem tamanho 1 acarretará em erro 
	fmt.Println(<-canal)
}