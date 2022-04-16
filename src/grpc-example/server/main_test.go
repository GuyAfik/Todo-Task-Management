package main

import (
	"testing"
)

func TestSomething(t *testing.T) {
	total := Sum(5, 5)
	t.Log(total)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
