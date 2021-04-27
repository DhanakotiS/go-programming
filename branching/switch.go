package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list := []int{10, 20, 30, 40, 50}

	// Switch
	n := rand.Intn(len(list))
	pick := list[n]
	switch pick {
	case 10:
		fmt.Println("Ten is Picked.")
	case 20:
		fmt.Println("Twenty is Picked.")
	case 30:
		fmt.Println("Thirty is Picked.")
	default:
		fmt.Println("Something else is picked.")
	}

	//Switch with short statements
	n = rand.Intn(len(list))
	switch pick = list[n]; pick {
	case 10:
		fmt.Println("Ten is Picked.")
	case 20:
		fmt.Println("Twenty is Picked.")
	case 30:
		fmt.Println("Thirty is Picked.")
	default:
		fmt.Println("Something else is picked.")
	}
}
