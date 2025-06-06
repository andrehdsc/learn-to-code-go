package main

import "fmt"


func main() {
	fmt.Println(x())
}

var x = func() int {
	return 42
}
