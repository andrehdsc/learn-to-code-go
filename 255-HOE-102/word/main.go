//Package word provides custom functions for working with strings.
package word

import "strings"

// UseCount returns a hashmap that contains the number of times a given word is used in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a given string. 
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
