package main

import "fmt"

func main() {
	// There are various primitive datatypes in go

	// 1. Booleans - bool
	// 2. Integers - int8, int16, int32, int64, int
	// 3. Unsigned integers - uint8, uint16, uint32	, uint64, uint, uintptr
	// 4. Floating point numbers - float32, float64
	// 5. Complex numbers - complex64, complex128
	// 6. Strings - string

	var boolean bool = true
	var integer int = 32
	var unSignedInteger uint = 32
	var floatNumber float32 = 32.32
	var complexNumber complex64 = 32 + 32i
	var str string = "Hello World"

	fmt.Println(boolean)
	fmt.Println(integer)
	fmt.Println(unSignedInteger)
	fmt.Println(floatNumber)
	fmt.Println(complexNumber)
	fmt.Println(str)
}
