package main

import (
	"fmt"

	"github.com/andrehdsc/learn-to-code-go/255-HOE-102/quote"
	"github.com/andrehdsc/learn-to-code-go/255-HOE-102/word"
)

func main() {
	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	fmt.Println(word.Count(quote.SunAlso))
}
