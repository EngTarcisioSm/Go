/*

*/
package main 

import(
	"fmt"
	"log"
)

func main(){
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//informa no retorno da mensagem de erro alguma informação referente a entrada que vausou o problema, ou seja, um error.new com sprintf() extra
		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
	}
	return 42, nil
}

/*PAREI EM 8:30/14:05 DA AULA CAP23 - TRATAMENTO DE ERROS - 5. ERROS COM INFORMAÇÕES ADICIONAIS*/