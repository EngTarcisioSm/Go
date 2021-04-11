/*
	- Partindo do código abaixo, ordene os []user por idade e sobrenome.
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenaPorIdade 	[]user
type ordenaPorSobrenome []user

func (x ordenaPorIdade) Len() int{return len(x)}
func (x ordenaPorIdade) Less(i, j int) bool {return x[i].Age < x[j].Age}
func (x ordenaPorIdade) Swap(i,j int) {x[i], x[j] = x[j], x[i]}

func (x ordenaPorSobrenome) Len() int{return len(x)}
func (x ordenaPorSobrenome) Less(i, j int) bool{return x[i].Last < x[j].Last}
func (x ordenaPorSobrenome) Swap(i,j int) {x[i], x[j] = x[j], x[i]}







func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	fmt.Println(".........................................")
	// your code goes here
	sort.Sort(ordenaPorIdade(users))
	for x:= 0; x<len(users);x++{
		fmt.Println(users[x])
	}
	fmt.Println(".........................................")
	sort.Sort(ordenaPorSobrenome(users))
	for x:= 0; x<len(users);x++{
		fmt.Println(users[x])
	}
}