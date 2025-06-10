package main

import "fmt"

type dude struct {
	first string
}

func main() {
	p1 := dude{"James"}
	fmt.Println("My name is ", p1.first)

	p1 = changeNameV(p1, "Jaime")
	fmt.Println("My name is ", p1.first)

	changeNameP(&p1, "Jimmy")
	fmt.Println("My name is ", p1.first)

}

func changeNameV(p dude, s string) dude {
	p.first = s
	return p
}

func changeNameP(p *dude, s string) {
	p.first = s
}
