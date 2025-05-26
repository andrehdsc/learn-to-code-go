package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := range 42 {
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Println("Iteration", i, "\tX = 0")
		case 1:
			fmt.Println("Iteration", i, "\tX = 1")
		case 2:
			fmt.Println("Iteration", i, "\tX = 2")
		case 3:
			fmt.Println("Iteration", i, "\tX = 3")
		case 4:
			fmt.Println("Iteration", i, "\tX = 4")
		}
	}
}
