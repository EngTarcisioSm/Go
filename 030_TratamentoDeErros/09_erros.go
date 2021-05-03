/*
- Erros com informações adicionais (erros custumizados)
	- Para que nossas funções retornem erros customizados, podemos utilizar:
		- return errors.New()
		- return fmt.Errorf(), possui um errors.New() embutido (https://golang.org/pkg/builtin/#error)
	- A interface error do package built-in é a interface convencional para representar uma condição de erro, onde um valor vazio representa "nenhum erro"
		- Para implementar a interface error é so necessário criar um tipo que tenha o método Error() retornando uma string 
	- Em go valores de erros não são especiais, eles são valores como quaisquer outros, é possivel usar a linguagem inteira para tratar desses erros do mesmo jeito que qualquer outro valor seria gerenciado 
	- Exemplos:
		- (10_erros.go) errors.New
		- (11_erros.go) var errors.New
		- (12_erros.go) fmt.Errorf
		- (13_erros.go) var fmt.Errorf
		- (14_erros.go) type + method = error interface  
*/
package main


func main(){

}