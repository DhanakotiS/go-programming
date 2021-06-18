package main

import "fmt"

type Stu struct {
	name string
	age  int
}

func (s Stu) print() {
	fmt.Println(s.name, s.age)
}

func (s *Stu) change(name string) {
	s.name = name
	fmt.Println(*s)
}

func print(s Stu) {
	fmt.Println(s.name, s.age)
}

func change(s *Stu, name string, age int) {
	s.name = name
	(*s).age = age
	fmt.Println(*s)
}

func main() {
	s := Stu{"John", 21}
	s.print()

	s = Stu{"John", 22}
	s.change("Doe")

	s = Stu{"Will", 25}
	print(s)

	s = Stu{"Steve", 25}
	change(&s, "Wan", 22)
}
