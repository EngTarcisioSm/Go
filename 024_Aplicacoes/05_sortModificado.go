/*
- O sort que eu quero não existe. Quero fazer o meu.
- Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
    - type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
- Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
- E aí posso fazer do jeito que eu quiser.
- Exemplo:
    - struct carros: nome, consumo, potencia
    - slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
    - tipo ordenarPorPotencia
    - tipo ordenarPorConsumo
    - tipo ordenarPorLucroProDonoDoPosto



	Customização de um sort em função de uma struct, utilizando a logica de interface  sort organiza qualquer interface do tipo Interface. Interface implementa len() Less() e Swap() caso exista essas 3 implementações é possivel implementar essas funções em estruturas tornando-as do tipo Interface podendo utilizar sort
*/

package main

import ("fmt"
		"sort"
)

type carro struct {
	nome 		string 
	potencia 	int
	consumo 	int
}

type ordenarPorPotencia []carro
type ordenarPorConsumo []carro
type ordenarPorLucroProDonoDoPosto []carro

func(a ordenarPorPotencia) 	Len() int{ return len(a) }
func(a ordenarPorPotencia)	Less(i, j int) bool{ return a[i].potencia < a[j].potencia}
func(a ordenarPorPotencia)	Swap(i,j int) {a[i], a[j] = a[j],a[i]}

func(a ordenarPorConsumo) 	Len() int{ return len(a) }
func(a ordenarPorConsumo)	Less(i, j int) bool{ return a[i].consumo < a[j].consumo}
func(a ordenarPorConsumo)	Swap(i,j int) {a[i], a[j] = a[j],a[i]}

func(a ordenarPorLucroProDonoDoPosto) 	Len() int{ return len(a) }
func(a ordenarPorLucroProDonoDoPosto)	Less(i, j int) bool{ return a[i].consumo > a[j].consumo}
func(a ordenarPorLucroProDonoDoPosto)	Swap(i,j int) {a[i], a[j] = a[j],a[i]}



func main(){

	carros := []carro{
		carro{"Chevet",50,5},
		carro{"Porshe",300,5},
		carro{"Fusca",20,30},
	}

	fmt.Println("Sem ordenação:\n",carros)
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println("Ordenar por Potência:\n",carros)
	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println("Ordenar por Consumo:\n",carros)
	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println("Ordenar por Lucro Dono do Posto:\n",carros)
}