package main

import "fmt"

func main() {
	x := 127

	fmt.Printf("%d %x %#x\n", x, x, x)

	s := fmt.Sprintf("%d %x %#x", x, x, x)
	fmt.Println(s)
}
