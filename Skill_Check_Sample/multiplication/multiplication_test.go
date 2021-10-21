package main

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

// Test
func Testadd(t *testing.T) {
	a := add(4, 3)
	want := 7
	if a != want {
		t.Errorf("add() = %v, want %v", a, want)
	}
}
