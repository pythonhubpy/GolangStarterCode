package main

import "fmt"

func main() {
	// Variable declaration can be done in three different ways:

	// 1. Using the := operator
	integer_new := 32

	// 2. Using the var keyword
	var integer int

	// 3. Assigning value along with the var keyword
	var newInteger int = 42

	// Declaring multiple variables at once

	var (
		name string = "John"
		age  int    = 30
	)


	fmt.Println(integer)
	fmt.Println(integer_new)
	fmt.Println(newInteger)
	fmt.Println(name)
	fmt.Println(age)
}