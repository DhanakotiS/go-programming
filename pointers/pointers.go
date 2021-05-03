package main

import "fmt"

func main() {
	var a int = 36

	var ptr *int = &a

	fmt.Println(a)
	fmt.Println(ptr)

	*ptr = *ptr + 2

	fmt.Println(*ptr)
	fmt.Println(&a)

}
