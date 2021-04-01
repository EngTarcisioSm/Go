/*
	- Métodos
		- define a base de go para entender o motivo de tipos na linguagem go ser tudo

		- Método é uma função anexada a um tipo

		- Este método faz referencia a função que tinha como molde
			func (reciever) identificador(parametros)(returns){code}
		- Reciever funcionará como um parametro exclusivo daquela função,e e essa função só poderá ser utilizada na presença daquele tipo

		-O método adiciona funcionalidades a um tipo especifico atravez da criação desse método com o uso do parametro reciever de uma função
*/
package main

import "fmt"

/*criado um tipo pessoa*/
type pessoa struct{
	pessoa	string
	idade	int
}

/*o reciever dessa função é o tipo pessoa, ela só pode ser utilizada mediante o tipo pessoa*/
/*essa função não esta disponivel para qualquer uso, sendo possivel de ser usado apenas em conjunto com o tipo definido em reciever*/
func (p pessoa) oiBomDia(){
	fmt.Printf("%s de %v anos diz bom dia", p.pessoa,p.idade)
}


func main(){
	pessoa1 := pessoa{"Tarcisio",32}

	/*a função bom dia só pode ser utilizada mediante o uso de um tipo pessoa*/
	pessoa1.oiBomDia()

}