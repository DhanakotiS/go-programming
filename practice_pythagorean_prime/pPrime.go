package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	if isPrime(n) && n%4 == 1 {
		fmt.Println("TRUE")
	} else {
		fmt.Println(("FALSE"))
	}
}
