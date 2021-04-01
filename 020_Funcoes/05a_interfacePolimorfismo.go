/*
	Exemplo de uma implementação de interface e polimorfismo
*/
package main

import (
	"fmt"
	"math"
)

/*
	Interface geometria, para que faça parte dela, deve-se ter area() e perimetro()
*/
type geometria interface{
	area() 		float64
	perimetro() float64
}

/*tipo retangulo dois valores em sua estrutura*/
type retangulo struct{
	largura, altura	float64
}

/*Método utilizado apenas pelo tipo retangulo*/
func (r retangulo) area() float64{
	return r.altura * r.largura
}

/*metodo utilizado apenas pelo tipo retangulo*/
func (r retangulo) perimetro() float64{
	return 2*r.altura + 2*r.largura
}



/*tipo circulo tem um parametro*/
type circulo struct{
	raio float64
}

/*o tipo ciculo implementa o metodo area*/
func (c circulo) area() float64{
	return math.Pi * c.raio * c.raio
}

/*o tipo circulo implementa o metodo perimetro*/
func (c circulo) perimetro() float64{
	return 2 * math.Pi * c.raio
}

/*metodo da interface geometria*/
func medicao(g geometria){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

/*como circulo e retangulo implementam os metodos area e perimetro eles tambem possuem o typo geometria pois ele implementam os metodos necessários*/

func main(){
	r := retangulo{
		altura: 3, 
		largura: 4,
	}
	c := circulo{
		raio: 5,
	}

	medicao(r)
	medicao(c)
}

