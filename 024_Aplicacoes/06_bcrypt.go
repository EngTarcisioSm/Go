/*
 	- quando um package tem um "x" quer dizer que o mesmo esta em fase experimental
	 instalar o package com o comando
 	go get -u golang.org/x/crypto/bcrypt

	NÃO FUNCIONOU AINDA
*/

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "15011989"

	slicebyte, err := bcrypt.GenerateFromPassword([]byte(senha), 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", slicebyte)
	}
}
