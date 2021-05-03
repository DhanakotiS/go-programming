package main

import (
	"fmt"
	"time"
)

func pool(id int, send <-chan int, receive chan<- string) {
	for task := range send {
		receive <- fmt.Sprintln("Worker", id, "Started task", task)
		time.Sleep(10 * time.Millisecond)
		receive <- fmt.Sprintln("Worker", id, "Finished task", task)
	}
	close(receive)
}

func main() {
	s := make(chan int, 5)
	r := make(chan string)

	for id := 1; id <= 3; id++ {
		go pool(id, s, r)
	}
	for i := 1; i <= 5; i++ {
		s <- i
	}
	close(s)
	for i := range r {
		fmt.Print(i)
	}

}
