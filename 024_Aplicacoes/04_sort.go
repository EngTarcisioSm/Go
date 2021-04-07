/*
	package que realiza ordenação de slices utilizados string e int 
*/
package main

import ("fmt"
		"sort"
)

func main(){
	sliceString := []string{"Tarcisio","Bryan","Janete","Alipio"}

	fmt.Println(sliceString)

	sort.Strings(sliceString)

	fmt.Println(sliceString)

	sliceInt := []int {32,29,5,59,57}

	fmt.Println(sliceInt)
	
	sort.Ints(sliceInt)

	fmt.Println(sliceInt)
}