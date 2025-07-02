package main

import (
	"fmt"
	"log"
)
// TESTING EXERCISE!!!

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{"50.2289 N", "99.4656 W", fmt.Errorf("Negative number!(%v)", f)}
	}
	return 42, nil
}
