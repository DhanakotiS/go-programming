package main

import "fmt"

type stack struct {
	items []int
}

func (s *stack) Push(item int) {
	s.items = append(s.items, item)
	fmt.Println("Pushed item:", item)
}

func (s *stack) Pop() {
	top := len(s.items) - 1
	if top <= 0 {
		fmt.Println("Stack Empty")
	}
	fmt.Println("Item popped :", s.items[top])
	s.items = s.items[0:top]
}

func (s stack) Display() {
	for i := range s.items {
		fmt.Println("[", s.items[len(s.items)-1-i], "]")
	}
}

func main() {
	var s stack
	var choice, item int

LOOP:
	for {
		fmt.Println("Enter the Option (1, 2, 3, 4):\t 1. Push 2.Pop 3.Display 4. Quit")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter the item to push:")
			fmt.Scanln(&item)
			s.Push(item)
		case 2:
			s.Pop()
		case 3:
			s.Display()
		case 4:
			break LOOP
		}
	}
}
