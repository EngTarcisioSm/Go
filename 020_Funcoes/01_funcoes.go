/*
	- Funções tem como objetivo: 
		- abstrair funcionalidades
		- reutilização de código  
		- Corpo basico de uma função (toda função segue a estrutura abaixo)
			func (receiver) identifier(parameters)(returns){code}
				- func : keyword
				- receiver:
				- identifier: nome da função

		- A diferença entre parâmetro e argumentos: 
			- Funções são definidas com parametros 
			- Quando uma função vai ser utilizada se passa argumentos 
		
	- Tudo em go é "pass by value" passado como valor, diferente do python que é passado como referencia 
*/

package main


import "fmt"


func main(){
	basica()
	fmt.Println(nome(2,3))
	fmt.Println(soma(1,2,3,4,5,6,7,8))

	fmt.Println(soma2("tarcisio",1,3,5,6,7,9,1))
}

/*exemplo de função*/
func basica(){
	fmt.Println("bom dia!")
}


/*exemplo de função*/
/*
	receiver será visto em métodos - 
	algo a contar é um parametro
	int é o tipo de retorno, caso existisse mais de um retorno deveria estar todos entre parentesis e separados por virgula
*/
/*cada parametro tem de ser especificado seu tipo, como os dois tipos são iguais podem ser feitos da forma abaixo, é informado o tipo de retorno da função que é um int */
func nome (x, y int) int {
	return x + y
}

/*entradas n definidas*/
/*a entrada sera dado como um array */
func soma(x...int)(int,int){
	soma := 0
	for _, v := range x{
		soma += v
	}
	return soma, len(x)
}
/*uma entrada variadica e uma comum */
/*se tiver uma entrada variadica ela deve ser a ultima*/
func soma2(nome string, x...int)(int,int,string){
	soma := 0
	for _, v := range x{
		soma += v
	}
	string := "bom dia " + nome
	
	return soma, len(x), string
}