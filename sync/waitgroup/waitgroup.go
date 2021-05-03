package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i < 11; i++ {
		wg.Add(1)
		go squareNum(i, i*i, &wg)
	}

	wg.Wait()

	fmt.Println("All goroutines done, exiting...")
}

func squareNum(idx, n int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("In goroutine number", idx, "and Square of n is :", n*n)

}
