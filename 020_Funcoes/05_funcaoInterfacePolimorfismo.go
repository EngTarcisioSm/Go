/*
	- Interfaces e polimorfismos
		- Em go valores podem ter mais de um tipo
		- Uma interface permite que um valor tenha mais de um tipo
		- Interface é um conjunto de métodos
		- Todo tipo criado que tiver um conjunto de métodos fará parte de uma interface, ganhando com isso o tipo da interface
		- Exemplo
			- cachorro tem os métodos
				- come
				- dorme
				- enche o saco
			- gato tem os métodos
				- come
				- dorme
				- enche o saco
			- como eles tem esses métodos fazem eles fazerem parte da interface animal
			- alem de uma variavel ter o tipo cachorro, quando ela tiver, ela também tera o tipo animal, o mesmo se aplica para o gato
		- Em outras linguagens de programação é necessário informar que um tipo implementa outro tipo no caso de go caso ele cotenha alguns metodos predeterminados, automaticamente ele fara parte daquela interface,
		- Tendo uma interface implantada é possivel implementar o polimorfismo

		- Polimorfismo: atraves de uma mesma ação é possivel trabalhar em coisas de tipos distintos
			- ao invez de implementar dois metodos diferentes, é possivel implementar um metodo apenas que atenda dois ou mais tipos

		- Implementação de Interface
			- declaração keyword identifier type -> type x interface
			- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface
			- Se um tipo possuir todos os metodos ( que no caso da interface{} pode ser nenhum) então esse tipo iplicitamente(automatico) implementa a interface
		- Todo tipo implementa a interface vazia e a interface vazia não implementa nenhum método
*/
package main

import "fmt"

type pessoa struct{
	nome		string
	sobrenome	string
	idade 		int
}

type dentista struct{
	pessoa
	cro		int
	salario	float64
}

type engenheiro struct{
	pessoa
	crea	int
	salario	float64
}

func (x dentista) oibomdia(){
	fmt.Printf("Meu nome é %s %s e digo bom dia\n", x.nome, x.sobrenome)
}

func (x engenheiro) oibomdia(){
	fmt.Printf("Meu nome é %s %s crea %d e digo bom dia\n", x.nome, x.sobrenome, x.crea)
}

/*interface gente para fazer parte dela deve ter o método oibomdia implementado*/
type gente interface{
	oibomdia()
}

func serhumano(g gente){
	g.oibomdia()
	switch g.(type){
		case dentista:
			fmt.Printf("Meu CRO é: %v\n",g.(dentista).cro)
		case engenheiro:
			fmt.Printf("Meu CREA é: %v\n",g.(engenheiro).crea)
	}
}

func main(){
	dentista1 := dentista{
		pessoa: pessoa{
			nome: "Fulano",
			sobrenome: "da Silva",
			idade: 50,
		},
		cro: 000001,
		salario: 6000.01,
	}

	engenheiro1 := engenheiro{
		pessoa: pessoa{
			nome: "Tarcisio",
			sobrenome: "Souza de Melo",
			idade: 32,
		},
		crea: 20001,
		salario: 3400.00,
	}
	dentista1.oibomdia()
	engenheiro1.oibomdia()

	/*como tanto engenheiro como dentista tem os pre-requisitos para serem definidos como gente, é possivel atravez da interface chamar o metodo oibondia
		- como tanto engenheiro como dentista possuem oibomdia() isso faz com ele também tenha o tipo gente 
	*/
	/*
		a função apresenta comportamentos diferentes para os tipos que entram pois eles são gente, entretanto possuem o outro tipo que provoca funcionalidades diferentes
	*/
	/*
	Aqui entra o polimorfismo, com a mesma chamada de função obter comportamentos distintos
	*/
	serhumano(dentista1)
	serhumano(engenheiro1)
}