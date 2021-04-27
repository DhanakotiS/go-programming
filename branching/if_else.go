package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// If
	m := rand.Int63n(100)
	if m > 10 {
		fmt.Println(" m is greater dhan 10", m)
	}

	// If - Else | If with short statements
	if n := int64(rand.Int31n(50)); n > 20 {
		fmt.Println(" n is greater than 20", n)
	} else {
		fmt.Println(" n is less than 20", n)
	}

	//If -Else If
	o := rand.Int31n(200)
	if o < 50 {
		fmt.Println(" o is less than 50", o)
	} else if o < 100 {
		fmt.Println(" o is less than 100 and greater than 50", o)
	} else {
		fmt.Println(" o is greater than 100", o)
	}

}
