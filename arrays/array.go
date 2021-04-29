package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("Before Initialization :", arr)
	arr[0] = 10
	arr[2] = 12
	arr[3] = 13
	fmt.Println("After Initialization :", arr)

	a := [5]int{21, 22, 23, 24, 25}
	fmt.Println("Implicit Declared Array :", a)

	b := [...]int{11, 22, 33, 44} //Let the compiler infer the length of array [...]
	fmt.Println("Auto Detect Length of Array:", b)

	var c = [3]int{7, 8, 9}
	fmt.Println("Declaration with var keyword:", c)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for i, j := range a {
		fmt.Printf("Index :%v, Element :%v\n", i, j)
	}
	var i int = 0
	for i < len(b) {
		fmt.Println(b[i])
		i++
	}

	/* Below Statements causes errors due to unmatched datatypes
	str := [...]int{12, int("5"), int(12.5)}
	fmt.Println(str) */
}
