package main

import (
	"testing"
)

// Require returnvalue to be 1
func Test_equalateral_triangle(t *testing.T) {
	value := FindTriangle("2", "2", "2")

	if value != 1 {
		t.Errorf("Return value was wrong. Got: %d, expected: %d", value, 1)
	}
}

// Require returnvalue to be 2
func Test_isosceles_triangle(t *testing.T) {
	value := FindTriangle("1", "2", "2")

	if value != 2 {
		t.Errorf("Return value was wrong. Got: %d, expected: %d", value, 2)
	}
}

// Require returnvalue to be 3
func Test_scalene_triangle(t *testing.T) {
	value := FindTriangle("1", "2", "3")

	if value != 3 {
		t.Errorf("Return value was wrong. Got: %d, expected: %d", value, 3)
	}
}
