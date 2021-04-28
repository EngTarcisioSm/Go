package main

import "fmt"

func main(){
	/*
		retorna a quantidade de bytes impressos e um erro, entretanto esse erro ninguem verifica, a analise de erro abaixo é meramente didatica
	*/
	n, err := fmt.Println("Hello")

	if err != nil{
		/*
		não é verificado pois acarretaria em um loop infinito pois para cada impressão de erro seria necessario uma nova verificação de erro para a mensagem de erro anterior 
		*/
		fmt.Println(err)
	}
	fmt.Println(n)

	/*
	quando se deseja saber a informação de numero de bytes impressos, excepcionalmente neste caso é utilizado o underscore
	*/
	n2, _ := fmt.Println("Again")
	fmt.Println(n2)
}