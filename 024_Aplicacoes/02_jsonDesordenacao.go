/*
	- Converter JSON para uma struct em go

	- é necessário para o processo dar para a função de conversão um slice de bytes que sera convertido em uma struct

	- site que auxilia JSON-to-Go
*/
package main

import (
	"encoding/json"
	"fmt"
)
/*
	O campo entre crases são tags, funciona como um conversor de label pois o atributo na struct pode ter um nome e no JSON aparecer com outro nome, servindo como uma orientação para conversão, tanto de struct para JSON como de JSON para struct 
*/
type Informacoes struct {
	Nome          string `json:"Nome"`
	Sobrenome     string `json:"Sobrenome"`
	Idade         int    `json:"Idade"`
	Profissao     string `json:"Profissao"`
	Contabancaria int    `json:"Contabancaria"`
}

func main(){
	slicebytes := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"espião","Contabancaria":10000000}`)

	var James Informacoes
	err := json.Unmarshal(slicebytes,&James)

	if err != nil {
		fmt.Println("Error...")
	} else {
		fmt.Printf("%T \n %v", James, James)
	}
}