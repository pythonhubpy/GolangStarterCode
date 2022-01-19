package main

import "fmt"

const (
	enumeratedConstant = iota
	enumeratedConstantTwo
) // Enumerated constants

func main() {
	const integerConstant int = 32 // typed constant
	const unTypedConstant = 100    // un typed constant

	fmt.Println(integerConstant)
	fmt.Println(unTypedConstant)
	fmt.Println(enumeratedConstant)
	fmt.Println(enumeratedConstantTwo)
}
