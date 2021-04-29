package main

import "fmt"

func main() {

	var m map[string]int
	fmt.Println(m)
	/* A nil map has no keys, nor can keys be added.

	The make function returns a map of the given type, initialized and ready for use.*/

	m = make(map[string]int)

	// m["area"] = 12
	// m["radius"] = 5

	// fmt.Println(m)

	var val int
	var key string
	for i := 0; i < 3; i++ {
		fmt.Println("Enter Key:")
		fmt.Scanln(&key)
		fmt.Println("Enter Value:")
		fmt.Scanln(&val)
		m[key] = val
	}
	fmt.Println(m)
	fmt.Println("Enter a Key to delete:")
	fmt.Scanln(&key)
	delete(m, key)
	fmt.Println("After deletion: ", m)
}
