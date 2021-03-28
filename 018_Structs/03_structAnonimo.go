/*
	É um struct que não é previamente definido, podendo ser utilizado apenas dentro daquela variavel em que ele foi descrito
*/
package main

import "fmt"

func main(){
	x := struct{
		nome	string
		idade	int
	}{
		nome: "Mario",
		idade: 50,
	}

	fmt.Println(x)
}