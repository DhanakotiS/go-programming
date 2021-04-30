package main

import (
	"fmt"
)

func func_call(arr ...int) {
	var sum int = 0
	for _, j := range arr {
		sum += j
	}
	fmt.Println("Sum of Numbers:", sum)
}
func main() {
	var n int
	fmt.Println("Enter N:")
	fmt.Scan(&n)
	fmt.Println("Enter", n, "Numbers:")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	func_call(arr...)
}
