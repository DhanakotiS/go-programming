package main

import "fmt"

type Person struct {
	name string
	city string
}

func (p Person) String() string {
	return fmt.Sprint(p.name, " lives in ", p.city, "\n")
}

func main() {
	p1 := Person{"John Doe", "Chicago"}
	p2 := Person{"William", "Washington D.C."}

	fmt.Println(p1, p2)

}
