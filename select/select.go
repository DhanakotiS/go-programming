package main

import (
	"fmt"
	"time"
)

func run(ch chan string) {
	ch <- "Keep Running!"
}

func abort(qu chan string) {
	qu <- "Stop!"
}

func main() {
	ch := make(chan string)
	qu := make(chan string)

	go run(ch)
	time.Sleep(10 * time.Millisecond)

	go abort(qu)
	time.Sleep(10 * time.Millisecond)

	//Running select for cases 'channel which is ready 1st' or 'random if both chennals are ready'

	select {
	case c := <-ch:
		fmt.Println(c)
	case q := <-qu:
		fmt.Println(q)
	default:
		fmt.Println("No orders Received!")
	}

	//Running using loop for two channels
	for i := 0; i < 2; i++ {
		select {
		case c := <-ch:
			fmt.Println(c)
		case q := <-qu:
			fmt.Println(q)
		default:
			fmt.Println("No orders Received!")
		}
	}

	//If no. of channels are unknown, can be run with endless loop, break and label.
Label:
	for {
		select {
		case c := <-ch:
			fmt.Println(c)
		case q := <-qu:
			fmt.Println(q)
		default:
			fmt.Println("No orders received!")
			break Label
		}
	}
}
