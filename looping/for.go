package main

import "fmt"

func main() {

	// Thre component loop
	fmt.Println("Loop - init; cond ;inc - Print Even numbers")
	for i := 0; i <= 10; i = i + 2 {
		fmt.Println(i)
	}

	// 'while' kind of loop
	fmt.Println("Loop - cond - ")
	var i int = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// 'for-each' kind loop using range.
	fmt.Println("Loop : from, to")
	tes := []string{"Hello", "world", "Go", "is", "fun"}
	for i, k := range tes {
		fmt.Println(i, tes[i], k)
	}

	// 'infinite' type loop
	i := 0 
	for {
		if(i>5){
			break
		}
		fmt.Println(i)
		i = i+1
	}
}
