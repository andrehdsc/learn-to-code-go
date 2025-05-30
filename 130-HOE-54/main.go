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

	m := map[string]person{
		me.last: me,
		p1.last: p1,
	}
	for _, v := range m {
		fmt.Println(v.first, v.last, "favorite flavors are:")
		for _, vv := range v.fav {
			fmt.Println(vv)
		}
	}

}
