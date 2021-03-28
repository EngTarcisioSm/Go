/*
	Range e deletando em maps
*/
package main

import "fmt"

func main(){
	listaCompras := map[int]string{
		1: "arroz",
		2: "feijão",
		3: "pão de queijo",
		4: "Salmão",
	}
	fmt.Println(listaCompras)

	for key, value := range listaCompras {
		/*Em maps não é garantido */
		fmt.Println(key, value)
	}

	/*para utilizar delete que é uma função, ela tem como parametros o mapa ao qual será executado a ação e em seguida o key a ser deletado*/
	delete(listaCompras, 4)

	fmt.Println()

	for key, value := range listaCompras {
		/*Em maps não é garantido */
		fmt.Println(key, value)
	}
}