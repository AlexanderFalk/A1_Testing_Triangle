package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func scalene_triangle() {
	fmt.Println("Your triangle is of type: scalene triangle!")
}

func equilateral_triangle() {
	fmt.Println("Your triangle is of type: equalateral triangle!")
}

func isosceles_triangle() {
	fmt.Println("Your triangle is of type: isosceles triangle!")
}

func FindTriangle(value1, value2, value3 string) int {
	var returnValue int
	if value1 == value2 && value2 == value3 {
		equilateral_triangle()
		returnValue = 1
		return returnValue
	}
	if (value1 == value2 && value1 != value3) || (value2 == value3 && value2 != value1) || (value1 == value3 && value1 != value2) {
		isosceles_triangle()
		returnValue = 2
		return returnValue
	}
	if value1 != value2 && value2 != value3 && value1 != value3 {
		scalene_triangle()
		returnValue = 3
		return returnValue
	}

	return 0
}

func main() {
	triangle_test := flag.NewFlagSet("whatis", flag.ExitOnError)
	values := triangle_test.String("values", "", "Add three values to test what type the triangle is")
	fmt.Println(*values)
	var arrayValues []string
	switch os.Args[1] {
	case "whatis":
		err := triangle_test.Parse(os.Args[2:])
		check(err)
		arrayValues = strings.Split(*values, " ")
		fmt.Println("First value: ", arrayValues[0])
		fmt.Println("Second value: ", arrayValues[1])
		fmt.Println("Third value: ", arrayValues[2])

	default:
		fmt.Printf("%s", "Error")
	}
	if triangle_test.Parsed() {
		fmt.Println("Values inserted...", os.Args[2:])
		FindTriangle(arrayValues[0], arrayValues[1], arrayValues[2])
		/*
			if arrayValues[0] == arrayValues[1] && arrayValues[1] == arrayValues[2] {
				equilateral_triangle()
				returnValue = 1
				return returnValue
			}
			if (arrayValues[0] == arrayValues[1] && arrayValues[0] != arrayValues[2]) || (arrayValues[1] == arrayValues[2] && arrayValues[1] != arrayValues[0]) || (arrayValues[0] == arrayValues[2] && arrayValues[0] != arrayValues[1]) {
				isosceles_triangle()
				returnValue = 2
				return returnValue
			}
			if arrayValues[0] != arrayValues[1] && arrayValues[1] != arrayValues[2] && arrayValues[0] != arrayValues[2] {
				scalene_triangle()
				returnValue = 3
				return returnValue
			}
		*/
	}
}
