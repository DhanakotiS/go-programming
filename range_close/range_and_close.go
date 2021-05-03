package main

import "fmt"

func squares(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i * i
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go squares(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}

}
