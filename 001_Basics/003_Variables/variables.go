package main

import "fmt"

var global = "Global var."

func main() {
	x := 99
	fmt.Println(x)

	y := x
	fmt.Println(y)

	x += 10
	fmt.Println(x, y)

	var s string = "123"
	fmt.Println(x, y, s, global)

	fmt.Printf("%d - %T\n", x, x)
	fmt.Printf("%s - %T\n", s, s)

	var changeling = `ASDF`
	fmt.Printf("%v - %T\n", changeling, changeling)

	// Does not work!
	// changeling = 1234
	// fmt.Printf("%s - %T\n", changeling, changeling)

	var character = '%'
	fmt.Printf("%c - %T\n", character, character)
}
