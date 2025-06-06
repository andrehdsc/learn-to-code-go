package main

import "fmt"

func main() {

	fmt.Println(func() int {
		return 42
	}())
}

//Create an annonymous func.
