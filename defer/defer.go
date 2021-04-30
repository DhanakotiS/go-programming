package main

import "fmt"

func main() {
	fmt.Println("This statement runs 1st")

	defer fmt.Println("This statement runs last : Defer")

	fmt.Println("This statement runs 2nd")

	defer fmt.Println("This statement runs 4th : Defer")

	fmt.Println("This statement runs 3rd")
}
