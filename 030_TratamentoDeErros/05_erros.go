/*
 	=> Cria um arquivo novo
	=> Verifica a existencia de erro
	=> Caso ocorra um erro o programa é encerrado
	=> Caso não ocorra erros é dado a instrução para o fechamento do arquivo com defer (para ser o ultimo a ser executado)
	=> É inserido no arquivo texto uma string
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	f, err := os.Create("names.txt")
	if err != nil{
		fmt.Println(err)
	}

	defer f.Close()
	defer fmt.Println("Rodou a função que estava em defer")

	r:=strings.NewReader("James Bond")
	io.Copy(f,r)

	/*parei em cap.23 - Tratamento de Erros - 2 Verificando erros*/
}