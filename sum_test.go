package main

import "testing"

func TestSum(t *testing.T) {
	a := 1
	b := 2
	if sum(a, b) != a+b {
		t.Fail()
	}
}
