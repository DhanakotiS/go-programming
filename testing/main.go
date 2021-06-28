package main

import "fmt"

func Add(lis []int) int {
	var sum int = 0
	for _, i := range lis {
		sum = sum + i
	}
	return sum
}

func main() {
	lis := []int{12, 11, 56, 45, 85, 1, 64}
	fmt.Println(Add(lis))
}
