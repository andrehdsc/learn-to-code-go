package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi! My name is %v and i'm %v years old. 😁", p.first, p.age)
	return
}

func main() {
	p := person{
		first: "John",
		age:   31,
	}
	p.speak()
}
