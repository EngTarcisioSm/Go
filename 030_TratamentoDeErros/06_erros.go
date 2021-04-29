/*leitura de um arquivo
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)


func main(){
	f, err := os.Open("names.txt")

	if err != nil{
		fmt.Println(err)
		return 
	}

	defer f.Close()
	/*
		- a função le todo o arquivo
		- é retornado uma cadeia de bytes
	*/ 
	bs, err := ioutil.ReadAll(f)

	if err != nil{
		fmt.Println(err)
		return 
	}

	fmt.Println(string(bs))
}