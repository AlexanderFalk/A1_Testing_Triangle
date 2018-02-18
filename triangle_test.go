package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Require returnvalue to be 1
func TestEquilateralTriangle(t *testing.T) {
	//Arrange & act
	value := FindTriangle("2", "2", "2")

	//assertion
	expectedValue := 1
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value, expectedValue)
}

// Require returnvalue to be 2
func TestIsoscelesTriangle(t *testing.T) {
	//Arrange & act
	value := FindTriangle("1", "2", "2")

	//assertion
	expectedValue := 2
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value, expectedValue)
}

// Require returnvalue to be 3
func TestScaleneTriangle(t *testing.T) {
	//Arrange & act
	value := FindTriangle("13", "9", "14")

	//assertion
	expectedValue := 3
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value, expectedValue)
}

// Require returnvalue to be 3
func TestInvalidTriangle(t *testing.T) {
	//Arrange & act
	value := FindTriangle("1", "2", "6")

	//assertion
	expectedValue := 4
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value)
}

//Added from 16-02-2018 for the purpose of assignment #2

//Should return 4 - Invalid
func TestTriangleOne(t *testing.T) {
	//Arrange & act
	value := FindTriangle("3", "2", "8")

	//assertion
	expectedValue := 4
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value, expectedValue)
}

//Should return 2 - Isosceles
func TestTriangleTwo(t *testing.T) {
	//Arrange & act
	value := FindTriangle("10", "8", "8")

	//assertion
	expectedValue := 2
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value, expectedValue)
}

//Should return 4 - Invalid
func TestTriangleThree(t *testing.T) {
	//Arrange & act
	value := FindTriangle("1", "2", "3")

	//assertion
	expectedValue := 4
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value, expectedValue)
}

//Checks for zero length values
func TestCheckZeroLengthValue(t *testing.T) {
	//Arrange & act
	value := FindTriangle("0", "2", "3")

	//assertion
	expectedValue := 0
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value, expectedValue)
}

//Should return 4 - Invalid
func TestCheckMinusValue(t *testing.T) {
	//Arrange & act
	value := FindTriangle("1", "2", "-2")

	//assertion
	expectedValue := 0
	assert.Equal(t, expectedValue, value, "Got: %d, expected: %d", value, expectedValue)
}
