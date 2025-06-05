package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(7, 5)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}
func TestSubtract(t *testing.T) {
	got := subtract(17, 5)
	want := 12
	if got != want {
		t.Errorf("Subtract was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(7, 5, add)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}
func TestDoMath2(t *testing.T) {
	got := doMath(17, 5, subtract)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}