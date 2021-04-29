package main

import "fmt"

func main() {

	var a []int
	a = append(a, 0)
	a = append(a, 2, 3, 5)
	fmt.Println(a)

	b := []int{5, 6, 4, 8, 9, 3, 2, 21, 45}
	fmt.Println(b)

	ln := 5
	cp := 10
	c := make([]int, ln, cp)
	// c = append(c, 12, 11, 10)
	// c = append(c, 9, 5, 3)
	c[0] = 10
	c = append(c, 1)

	fmt.Println(c)
	fmt.Println(c[0:10])
	fmt.Println(c[ln:cp])

	fmt.Println("A:", len(a), cap(a))
	fmt.Println("B:", len(b), cap(b))
	fmt.Println("C:", len(c), cap(c))
}
