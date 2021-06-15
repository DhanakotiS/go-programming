package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	l float64
	b float64
}
type circle struct {
	r float64
}

func (R *rectangle) area() float64 {
	return R.l * R.b
}

func (R *rectangle) perimeter() float64 {
	return 2 * (R.l + R.b)
}
func (C *circle) area() float64 {
	return math.Pi * C.r * C.r
}

func (C *circle) perimeter() float64 {
	return 2 * math.Pi * C.r
}
func main() {
	var l, b, r float64
	var R rectangle
	var C circle
	fmt.Scan(&l)
	fmt.Scan(&b)
	fmt.Scan(&r)
	R.l = l
	R.b = b
	C.r = r
	fmt.Println(R.area())
	fmt.Println(R.perimeter())
	fmt.Println(C.area())
	fmt.Println(C.perimeter())
}

// 5
// 6
// 5

// 30
// 22
// 78.53981633974483
// 31.41592653589793
