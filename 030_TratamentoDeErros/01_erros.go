/*
	-> Em go não temos exceções → https://golang.org/doc/faq#exceptions
		-> Golang acredita que juntar exceções ao programa e utilizar coisas como try-catch-finally resulta em código bagunçado e tratar situações como exceções quando não são, como por exemplo abrir arquivos 
		-> O golang possui uma tratativa diferente, utilizando multiplos retornos de funções, geralmente dentro desses multiplos retornos existe um valor de erro  indicando o sucesso ou não de determinada função. Caso o erro for vazio (nil) indica que ocorreu tudo bem, caso retorne algum valor indica um erro e procurar resolver essa situação 
		-> Existe um tipo erro (type error) 
		-> na opinião dos criadores da linguagem isso torna mais facil trabalhar com a linguagem 
		-> A linguagem go possui algumas funções para se recuperar de erros 
		-> Existem mecanismos na linguagem para tratar com exceções verdadeiras de fato 
			-> No go as exceções não são banalizadas
	-> o tipo erro é uma interface pré pronta da linguagem para representar erros, o valor vazio dentro deste tipo sinaliza ausencia de erros 
	-> só é necessário uma string para implementar o typo erro 
	-> package errors possui funções para manipular erros
		-> func New(text string) error, o text informado sera retornado em uma variavel do tipo error
	-> assim que se utiliza uma função que pode retornar erro na linha seguite ha efetue o tratamento para o possivel erro 
*/