package main

import "fmt"

func main() {
	fmt.Printf("Hello functions!\n")

	foo()
	bar()

	for i := 1; i <= 100; i++ {
		if i%10 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Printf("This is foo ...\n")
}

func bar() {
	fmt.Printf("This is bar ...\n")
}
