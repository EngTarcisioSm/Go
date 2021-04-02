/*
 	closure:
	 	- tradução para o portugues siginifca cercar ou capturar
		- Capturar um scope para que possamos utiliza-lo em outro contexto
		- Scopes conhecidos
			- Package level scope
			- Function level scope
			- Code block in code block scope
*/
package main

import "fmt"



func main(){
	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	/*outra variavel outro contexto*/
	b := i()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

/*
	- A função i() retorna uma outra função que tem retorno um int, dentro dessa função existe uma variavel que o contexto dela é salvo e pode ser utilizado a qualquer momento sem ser valor sofrer alguma alteração a não ser pela propria variavel que a armazenará. caso uma segunda variavel que utilize a função i(), será armazenado uma segunda variavel com o contexto salvo, modificavel apenas pela propria variavel 
*/
func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}