package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter()
}

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r *rectangle) area() float64 {
	return r.length * r.breadth
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *rectangle) perimeter() {
	fmt.Println(2 * r.length * r.breadth)
}
func (c circle) perimeter() {
	fmt.Println(2 * math.Pi * c.radius)
}
func main() {

	r := rectangle{12, 11}
	c := circle{5}

	s := []shape{&r, c}

	for _, sh := range s {
		fmt.Printf("%T\n", sh)
		fmt.Println("area\n", sh.area())
		fmt.Println("perimeter")
		sh.perimeter()
	}
}
