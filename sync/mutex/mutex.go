package main

import (
	"fmt"
	"sync"
)

var mutexVar sync.Mutex

func updateNumber(i int, n *int, wg *sync.WaitGroup) {
	defer mutexVar.Unlock()
	defer wg.Done()

	mutexVar.Lock()
	*n = *n * *n
	fmt.Println("n updates by goroutine ", i, " : n : ", *n)

}

func main() {
	var n int = 2
	var wg sync.WaitGroup
	fmt.Println("Base number :", n)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go updateNumber(i, &n, &wg)

	}
	wg.Wait()
}
