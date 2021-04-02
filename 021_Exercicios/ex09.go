/*
	- Callback: passe uma função como argumento a outra função.
*/
package main

import "fmt"


func main(){	
	printInvert(invert, "Tarcisio Souza de Melo")
}


func printInvert(f func(string) string,s string){
	str := f(s);
	fmt.Println(str)
}

func invert(s string) string{
	var str1 string
	x := len(s) -1
	for x >= 0 {
		str1 += string(s[x])
		x--
	}
	return str1
}
