package main

import "fmt"

//No Arguments no return
func greet() {
	fmt.Println("Hello, World!")
}

// Arguments but no return
func add(a int, b int) {
	fmt.Println(a + b)
}

// No Arguments but multiple return
func bio() (string, int) {
	var name string
	var age int
	fmt.Println("Enter Name:")
	fmt.Scanln(&name)
	fmt.Println("Enter Age:")
	fmt.Scanln(&age)
	return name, age
}

func main() {
	greet()
	add(5, 6)
	res := mod(12, 5)
	fmt.Println(res)
	name, age := bio()
	fmt.Println("Name :", name, "Age :", age)

}

//Arguments and return
func mod(x int, y int) int {
	return x % y
}
