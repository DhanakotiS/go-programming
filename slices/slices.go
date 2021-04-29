package main

import (
	"fmt"
)

func main() {
	var n, v int
	var odd, even = 0, 0
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		arr[i] = v
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			odd += arr[i]
		} else {
			even += arr[i]
		}
	}
	fmt.Println(odd)
	fmt.Println(even)
	fmt.Println(odd - even)
}
