package main

import (
	"fmt"
)

type indv struct {
}

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "julia",
		friends:   map[string]int{"Jhon": 21, "Jeff": 20, "Mary": 22},
		favDrinks: []string{"Scotch", "Coca-cola"},
	}

	fmt.Println("This person is", p1.first)
	fmt.Println(p1.first, "'s friends are:")
	for k, v := range p1.friends {
		fmt.Printf("%v age %v\n", k, v)
	}

	fmt.Println(p1.first, "'s favorite drinks are:")
	for _, v := range p1.favDrinks {
		fmt.Printf("%v\n", v)
	}
}
