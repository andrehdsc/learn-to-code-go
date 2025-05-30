package main

import (
	"fmt"
)

type car struct {
	engine
	make  string
	model string
	door  int
	color string
}

type engine struct {
	electric bool
}

func main() {
	c1 := car{
		engine: engine{electric: false},
		make:  "VW",
		model: "Gol",
		door:  5,
		color: "Silver",
	}
	c2 := car{
		engine: engine{electric: true},
		make:  "Chevy",
		model: "Bolt",
		door:  5,
		color: "White",
	}

	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println(c1.model)
	fmt.Println(c2.electric)
}
