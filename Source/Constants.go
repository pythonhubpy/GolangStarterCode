package main

import "fmt"

const (
	enumeratedConstant = iota
	enumeratedConstantTwo
)

func main() {
	const integerConstant int = 32
	const unTypedConstant = 100

	fmt.Println(integerConstant)
	fmt.Println(unTypedConstant)
	fmt.Println(enumeratedConstant)
	fmt.Println(enumeratedConstantTwo)
}