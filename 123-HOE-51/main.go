package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string)
	m["james_bond"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	m["no_dr"] = []string{"cats", "ice cream", "sunset"}
	m[`fleming_ian`] = []string{"steaks", "cigars", "spionage"}

	fmt.Println("Deleting...")
	delete(m, "no_dr")

	for kmap := range m {
		fmt.Println(kmap)
		for i, vx := range m[kmap] {
			fmt.Printf("%v:%v\t", i, vx)
		}
		fmt.Println()
		fmt.Println()
	}
}
