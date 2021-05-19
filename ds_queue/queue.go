package main

import "fmt"

type queue struct {
	items []int
}

func (q *queue) Insert(item int) {
	q.items = append(q.items, item)
}

func (q *queue) Delete() {
	if len(q.items) > 0 {
		top := q.items[0]
		fmt.Println("Item Deleted :", top)
		q.items = q.items[1:len(q.items)]
	} else {
		fmt.Println("Queue is Empty!")
	}
}

func (q queue) Display() {
	fmt.Println(q.items)
}

func main() {
	var q queue
	var item, choice int

LOOP:
	for {
		fmt.Print("Enter the option (1,2,3,4):\t 1.Insert 2.Delete 3.Display 4.Quit :")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter the number to Insert:")
			fmt.Scan(&item)
			q.Insert(item)
		case 2:
			q.Delete()
		case 3:
			q.Display()
		case 4:
			break LOOP
		}
	}
}
