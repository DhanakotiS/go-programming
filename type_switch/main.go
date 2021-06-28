package main

import (
	"fmt"
)

func myFunc(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("The Given Value is a Integer", a.(int))
	case string:
		fmt.Println("The Given Value is a String", a.(string))
	case float64:
		fmt.Println("The Given Value is Float", a.(float64))
	default:
		fmt.Println("Expected types not received")
	}
}

func main() {
	myFunc(32.06)
	myFunc(12)
	myFunc("Hello")
}
