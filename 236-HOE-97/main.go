package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal(fmt.Errorf("Error Converting to JSON: %v\n", err))	
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a any) ([]byte, error) {
	bs, err := json.Marshal(a)
	return bs, err
}
