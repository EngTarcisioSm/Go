/*
	- Canais são um dos tópicos mais interessantes da linguagem go
	-forma de transmitir dados, informações entre goroutines
	- Canais são o Jeito Certo® de fazer sincronização e código concorrente.
	- Eles nos permitem trasmitir valores entre goroutines.
	- Servem pra coordenar, sincronizar, orquestrar, e buffering.

	- Canais bloqueiam:
    - Eles são como corredores em uma corrida de revezamento
    - Eles tem que "passar o bastão" de maneira sincronizada
    - Se um corredor tentar passar o bastão pro próximo, mas o próximo corredor não estiver lá...
    - Ou se um corredor ficar esperando receber o bastão, mas ninguem entregar...
    - ...não dá certo.

	- Colocar informação no canal e retirar informação do canal deve ser feito de forma recorrente
		- uma goroutine coloca informação para uma outra goroutine utiliza-la
*/
package main

import "fmt"

func main(){
	//criando um canal palavra chave chan seguido do tipo da informação que trafegará 
	canal := make(chan int)
	//inserindo valores ao canal 
	//canal <- 42

	//não funciona pois o canal foi inserido valor e esta tentando se utilizar na mesma goroutine, lembrando que main se trata de uma goroutine
	//fmt.Println(<-canal)
	go func(){
		canal<- 43
	}()
	fmt.Println(<-canal)
}