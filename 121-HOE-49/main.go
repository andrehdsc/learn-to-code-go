package main

import (
	"fmt"
)

func main() {
	jlikes := []string{"shaken, not stirred", "martinis", "fast cars"}
	mlikes := []string{"intelligence", "literature", "computer science"}
	dlikes := []string{"cats", "ice cream", "sunset"}

	m := make(map[string][]string)
	m["james_bond"] = jlikes
	m["moneypenny_jenny"] = mlikes
	m["no_dr"] = dlikes

	for kmap := range m {
		fmt.Println(kmap)
		for i, vx := range m[kmap] {
			fmt.Printf("%v:%v\t", i, vx)
		}
		fmt.Println()
		fmt.Println()
	}
}
