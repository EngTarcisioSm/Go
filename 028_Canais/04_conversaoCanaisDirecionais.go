/*
	convertendo canais
*/
package main

import "fmt"

func main() {
	/*canal bidirecional*/
	c := make(chan int)
	/*canal recieve*/
	cr := make(<-chan int)
	/*canal send*/
	cs := make(chan<- int)

	/*
		Escrevendo o tipo dos canais
	*/
	fmt.Println("..........................")
	fmt.Printf("c\t% T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	/*conversões*/
	fmt.Println("..........................")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	/*não é possivel atribuir um tipo recieve a um send e vice versa*/
	//cs = cr

	/*
		conversão de geral para especifico é possivel
		conversão de especifico para geral não é possivel
		conversão de especifico para outro especifico não ocorre

		não é possivel atribuir tipos diferentes
	*/

}
