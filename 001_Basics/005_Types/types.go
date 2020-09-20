package main

import "fmt"

var a int

type testType int

var b testType

func main() {
	a, b = 19, 20

	fmt.Printf("%d - %T\n", a, a)
	fmt.Printf("%d - %T\n", b, b)

	// Conversion
	b = testType(a)
	fmt.Printf("%d - %T\n", b, b)
}
