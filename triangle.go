package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func scaleneTriangle() {
	fmt.Println("Your triangle is of type: scalene triangle!")
}

func equilateralTriangle() {
	fmt.Println("Your triangle is of type: equilateral triangle!")
}

func isoscelesTriangle() {
	fmt.Println("Your triangle is of type: isosceles triangle!")
}

func invalid() {
	fmt.Println("That is not a triangle my dear friend...")
}

//FindTriangle takes three parameters, which is the length of each side of a triangle
//It will then determine whether the triangle is:
//Scalene - No sides are equal
//Equalateral - All sides are equal
//Isosceles - Atleast two sides are equal
//The function also checks for invalid triangles
func FindTriangle(value1, value2, value3 string) int {

	a, err := strconv.Atoi(value1)
	check(err)
	b, err := strconv.Atoi(value2)
	check(err)
	c, err := strconv.Atoi(value3)
	check(err)

	//Used to check specific triangle test cases
	var returnValue int

	// Checks whether a triangle is valid or not. For a Triangle to be valid, you'll need the following formula:
	// a + b > c
	// a + c > b
	// b + c > a
	// If all is true, then you've a valid triangle
	if (a+b) > c && (a+c) > b && (b+c) > a {

		if value1 == value2 && value2 == value3 {
			equilateralTriangle()
			returnValue = 1
			return returnValue
		}
		if (value1 == value2 && value1 != value3) || (value2 == value3 && value2 != value1) || (value1 == value3 && value1 != value2) {
			isoscelesTriangle()
			returnValue = 2
			return returnValue
		}

		if value1 != value2 && value2 != value3 && value1 != value3 {
			scaleneTriangle()
			returnValue = 3
			return returnValue
		}
	} else {
		invalid()
		returnValue = 4
		return returnValue
	}
	return 0
}

func main() {
	triangleTest := flag.NewFlagSet("find", flag.ExitOnError)
	values := triangleTest.String("values", "", "Add three values to test what type the triangle is")
	fmt.Println(*values)
	var arrayValues []string
	switch os.Args[1] {
	case "find":
		err := triangleTest.Parse(os.Args[2:])
		check(err)
		arrayValues = strings.Split(*values, " ")
		fmt.Println("First value: ", arrayValues[0])
		fmt.Println("Second value: ", arrayValues[1])
		fmt.Println("Third value: ", arrayValues[2])

	default:
		fmt.Printf("%s", "Error")
	}
	if triangleTest.Parsed() {
		fmt.Println("Values inserted...", os.Args[2:])
		FindTriangle(arrayValues[0], arrayValues[1], arrayValues[2])
	}
}
