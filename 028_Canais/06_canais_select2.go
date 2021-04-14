/*
	- Segundo exemplo:
		- Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
		- Func 2 for infinito, select: case envia pra canal, case recebe de quit
	- Select torna possivel tanto enviar como receber
*/
package main

import "fmt"

func main(){
	 canal := make(chan int)
	 quit := make(chan int)

	 go recebeQuit(canal, quit)
	 enviaParaCanal(canal, quit)
}

func recebeQuit (canal chan int, quit chan int){
	for i := 0; i < 50; i++{
		fmt.Println("recebido: ", <-canal)
	}
	quit <- 0
}

func enviaParaCanal(canal chan int, quit chan int){
	qualquercoisa := 1
	for{
		select{
			case canal <- qualquercoisa:
				qualquercoisa++
				fmt.Println("virou ", qualquercoisa)
			case <-quit:
				return 
				
		}
	}
}