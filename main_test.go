package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	if 1 > 2 {
		t.Errorf("Maths is broken")
	}
}
