/*
- Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build nomeD0Arquiv.go
    - go install
		/apenas go install e ele instalara o cídgo no binario do sistema tornando possivel usa-lo apenas o chamando pelo nome sem extensão go algumaCoisaInstalada
*/

package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}