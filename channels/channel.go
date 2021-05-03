package main

import "fmt"

func say(name string, ch chan string) {
	s1 := fmt.Sprint("Hello ", name)
	ch <- s1
}

func display(grt string, ch chan string) {
	s2 := fmt.Sprint(grt, " Everybody")
	ch <- s2
}

func greet(str string) {
	fmt.Println(str)
}

func main() {
	ch := make(chan string)

	go say("John", ch)
	go display("Greetings", ch)

	msg, grt, test := <-ch, <-ch, "Test"

	greet("Hello, World")

	fmt.Println(msg)
	fmt.Println(grt)
	fmt.Println(test)

	/*
		Hello, World
		Greetings Everybody
		Hello John
		Test
	*/

	ch1 := make(chan string)
	ch2 := make(chan string)

	go say("John", ch2)
	go display("Greetings", ch1)

	msg, grt, test = <-ch2, <-ch1, "Test"

	greet("Hello, World")

	fmt.Println(msg)
	fmt.Println(grt)
	fmt.Println(test)

	/*
		Hello, World
		Hello John
		Greetings Everybody
		Test
	*/

}