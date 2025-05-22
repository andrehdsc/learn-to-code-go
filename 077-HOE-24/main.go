package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Println("x: ", x)

	switch {
	case x > 0 && x <= 100:
		fmt.Println("between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("between 101 and 200")
	case x > 200 && x <= 250:
		fmt.Println("between 201 and 250")
	}

	// if x > 0 && x <= 100 {
	// 	fmt.Println("between 0 and 100")
	// } else if x > 100 && x <= 200 {
	// 	fmt.Println("between 101 and 200")
	// } else {
	// 	fmt.Println("between 201 and 250")
	// }
}

func init() {
	fmt.Println("This is where initialization for my program occurs.")
}

//niladic function means no arguments
