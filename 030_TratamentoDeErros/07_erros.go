/*
	=> forma de lidar com o erro 01
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	_, err := os.Open("no-file.txt")
	if err != nil{
		/*
			=> printa na tela a mensagem de erro na forma abaixo
			err happend open no-file.txt: no such file or directory
		*/
		fmt.Println("err happend", err)

		/*
			=> printa na tela a mensagem de erro na forma abaixo
			insere data e hota
		*/
		log.Println("err happend",err)
		/*
			=> printa na tela a mensagem de erro na forma abaixo
			E ENCERRA O PROGRAMA 
		*/
		//log.Fatalln(err)
		/*
			=> printa na tela a mensagem de erro na forma abaixo
		*/
		log.Panicln("err")

		 
	}
}