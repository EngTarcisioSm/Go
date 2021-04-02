/*
	Função que retorna uma função
*/

package main

import "fmt"

/*
	A função retfun não possui parametros entretanto returna uma função que tem parametro int que retorna um int
*/
func retfunc() func(int) int{
	return func(i int) int{
		return i * 10
	}
}

func main(){
	x := retfunc()

	y := x(2)

	fmt.Println(y)

	//ou

	z := retfunc()(5)
	fmt.Println(z)

}