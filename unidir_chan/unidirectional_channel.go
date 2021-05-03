package main

import (
	"fmt"
	"time"
)

func put(ch chan<- string) {
	ch <- "Hello, World!"
	ch <- "Bonjour le monde!"

	// Below line causes error because cannot receive from send-only channel.
	// french := <-ch
}

func get(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)

	// Below line causes error : cannot send to receive only channel
	// ch <- "Ciao mondo!"
}

func main() {
	ch := make(chan string)

	go put(ch)

	go get(ch)

	time.Sleep(10 * time.Millisecond)

	fmt.Println(<-ch)
}
