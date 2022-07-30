package main

import (
	"fmt"
)

func main() {
	// Declare a variable of type string with a value of "Hello World"
	var str string = "Hello World"
	// Declare a variable of type bool implicitly with a value of true
	bool := true
	// Declare a variable of type int with a value of 10
	var int int = 10
	// Declare a variable of type float64 with a value of 10.5
	var float64 float64 = 10.5
	// Declare a variable of type complex128 with a value of 1 + 2i
	var complex128 complex128 = 1 + 2i
	// Declare a variable of type rune with a value of 'a'
	var rune rune = 'a'
	// Declare a variable of type string with a value of "Hello World"
	fmt.Println(str, bool, int, float64, complex128, rune)
	a, b := 20, 30
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
