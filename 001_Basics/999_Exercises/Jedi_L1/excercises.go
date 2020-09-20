package main

import "fmt"

func main() {
	exercise1()
	fmt.Println()

	exercise2()
	fmt.Println()

	exercise3()
	fmt.Println()

	exercises4and5()
}

func exercise1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exercise2() {
	var x int
	var y string
	var z bool

	fmt.Println(x, "'"+y+"'", z)
}

var a int = 42
var b string = "James Bond"
var c bool = true

func exercise3() {
	var buffer string = fmt.Sprintf("%d %s %t", a, b, c)
	fmt.Println(buffer)

	s := fmt.Sprintf("%T", buffer)
	fmt.Println(s)
}

type myIntType int

var myGlobalVar int

func exercises4and5() {
	var myVar myIntType

	fmt.Printf("myVar: %v, type of myVar: %T\n", myVar, myVar)

	myVar = 42
	fmt.Println(myVar)

	myGlobalVar = int(myVar)
	fmt.Printf("value: %v, type: %T\n", myGlobalVar, myGlobalVar)
}
