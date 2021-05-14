package main

import (
	"fmt"
	"math"
	"strconv"
)

func isNarcissistic(num int, n int) {
	var sum, count float64 = 0, 0
	var val float64 = float64(num)
	var temp int
	for {
		temp = num % 10
		num = num / 10
		sum += math.Pow(float64(temp), float64(n))
		if num == 0 {
			break
		}
		count += 1
	}
	if sum == val {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
func main() {
	var num, n int
	fmt.Scan(&num)
	s := strconv.FormatInt(int64(num), 10)
	n = len(s)

	isNarcissistic(num, n)

}
