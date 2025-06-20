package main

import (
	"fmt"
)

type Person struct {
	frase string
}

func (p *Person) Speak() {
	fmt.Println(p.frase)
}

type human interface {
	Speak()
}

func SaySomething(h human) {
	h.Speak()
}

func main() {

	p := Person{"I am Iron Man."}
	SaySomething(&p)

}
