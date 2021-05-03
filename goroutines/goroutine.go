package main

import (
	"fmt"
	"time"
)

func printer(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(10 * time.Millisecond)
	}
}

func printName(n string) {
	for i := 0; i < 5; i++ {
		fmt.Println(n)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go printName("John")
	printer("Hello, World!")
}

/*
Hello, World!
John
John
Hello, World!
Hello, World!
John
*/
