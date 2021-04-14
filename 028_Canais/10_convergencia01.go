/*
	- Convergencia
		- Padrão de comunicação, uma forma de se fazer as coisas;
		- Pegar dados de vários pontos e convergir eles para apenas alguns poucos ou um unico

*/
/*
	- Exemplo:
	- Canais par, ímpar, e converge.
        - Func send manda pares pra um, ímpares pro outro, depois fecha.
        - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par
*/
package main

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

}

func envia(p, i chan int) {
	x := 100
	for n := 1; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
}

//parei em 2:45
