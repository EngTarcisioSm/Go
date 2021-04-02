/*
	Callback nada mais é do que passar uma função como argumento de outra função
*/

package main

import "fmt"

func main(){
	t := somentepares(soma, []int{50,10,40,28,32}...)

	fmt.Println(t)

	u := somenteImpares(soma, []int{1,2,3,4,5,6,7,8,9}...)

	fmt.Println(u)

}

func soma(x...int) int{
	n := 0
	for _,v := range x {
		n += v
	}
	return n 
}

/*A função abaixo recebe uma função como parametro essa função possui como parametro um slice de int e retorna um int, o segundo parametro é um slice de inteiros*/
func somentepares(f func(x...int)int,y...int)int{
	var slice []int

	for _,v := range y{
		if v % 2 == 0{
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}

func somenteImpares(f func(x...int)int,y...int)int{
	var slice2 []int

	for _,v := range y{
		if v % 2 != 0{
			slice2 = append(slice2, v)
		}
	}
	total := f(slice2...)
	return total
}