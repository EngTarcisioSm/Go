/*
	- Canais podem ser direcionais.
		- canais podem ser configurados apenas para receber informações ou apenas para enviar
	- E isso serve pra...?
	- Um send channel e um receive channel são tipos diferentes. Isso permite que os type-checking mechanisms do compilador façam com que não seja possível, por exemplo, escrever num canal de leitura.
	- Aos aventureitos: https://stackoverflow.com/questions/1...​
	- Canais bidirecionals (send & receive)
		- send chan←
			- erro apresentado caso se deseje tirar informação do mesmo
			- error: "invalid operation: ←cs (receive from send-only type chan← int)"
		- receive ←chan
			- erro apresentado caso se deseje receber informação do mesmo
			- error: "invalid operation: cr ← 42 (send to receive-only type ←chan int)"
	- Exemplo: https://play.golang.org/p/TlcSm8bHkW​
	- A seta sempre aponta para a esquerda.
	- Assignment/conversion: 
		- de geral para específico
		- de específico para geral não
		- Exemplos:
			- geral pra específico: https://play.golang.org/p/H1uk4YGMBB​
			- específico pra específico: https://play.golang.org/p/8JkOnEi7-a​ 
			- específico pra geral: https://play.golang.org/p/4sOKuQRHq7​
			- atribuição tipos !=: https://play.golang.org/p/bG7H6l03VQ​ 

	Utilizando canais

	- Em funcs podemos especificar:
		- receive channel
			- Parâmetro receive channel: (c ←chan int)
			- No scope dessa função, esse canal só recebe
			- Não podemos fechar um receive channel
		- send channel 
			- Parâmetro send channel: (c chan← int)
			- No scope dessa função, esse canal só envia
			- Podemos fechar um send channel
*/

package main

import "fmt"

func main(){
	/*
		- O canal é inicializado como bidirecional, entretanto poderia ser configurado de inicio como apenas de leitura ou apens de escrita
		make(<-chan int)
		make(->chan int)
	*/
	canal := make(chan int)

	/*
		ao fazer isso o canal de proposito geral vira um canal send internamente a função send
	*/
	go send(canal)
	/*
		O canal é convertido para um recieve
	*/
	recieve(canal)
}
/*
	a função receberá como parametro um canal de inserção de inteiros 
*/
func send(s chan<- int){
	s <- 42
}

func recieve(r <-chan int){
	/*<- indica que esta sendo tirado informação do canal */
	fmt.Println("O valor recebido do canal foi ", <-r)
}