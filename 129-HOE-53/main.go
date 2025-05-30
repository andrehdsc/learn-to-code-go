package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	me := person{
		first: "Andre",
		last:  "Casalli",
		fav:   []string{"Chocolate", "Vanilla", "Strawberry"},
	}
	p1 := person{
		first: "Jhon",
		last:  "Doe",
		fav:   []string{"Banana", "Mint", "Vanilla"},
	}
	fmt.Printf("%v %v:\n", me.first, me.last)
	for _, p := range me.fav {
		fmt.Println(p)
	}
	fmt.Printf("%v %v:\n", p1.first, p1.last)
	for _, p := range p1.fav {
		fmt.Println(p)
	}
}
