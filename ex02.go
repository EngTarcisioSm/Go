/*
 Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/
package main

import "fmt"

type humanos interface{
	falar()
}

type pessoa struct{
	nome	  	string
	sobrenome 	string
	idade		int64
}

func (p *pessoa) falar(){
	fmt.Println(p.nome, p.sobrenome,p.idade)
}

func dizerAlgumaCoisa(h humanos){
	h.falar()
}

func main(){
	tarcisio := pessoa{
		nome: "Tarcisio",
		sobrenome: "Souza de Melo",
		idade: 32,
	}

	dizerAlgumaCoisa(&tarcisio)
	/*
		- Observações:
			a interface foi implementada para o ponteiro de pessoa, quando utiliza-se dizerAlgumaCoisa(tarcisio)
			e não funciona pois a interface foi implementada apenas para o ponteiro, quando utilizamos 
			tarcisio.falar() funciona pois na linguagem go tarcisio = &tarcisio, quando utilizamos structs apesar
			de o correto ser utilizar dizerAlgumaCoisa(&tarcisio), (&tarcisio).falar()
	*/
}