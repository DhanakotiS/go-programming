package main

import "fmt"

type Marks struct {
	name string
	sub1 int
	sub2 int
	sub3 int
	sub4 int
}

func (m Marks) average() int {
	return ((m.sub1 + m.sub2 + m.sub3 + m.sub4) / 4)
}
func (m *Marks) change() {
	m.sub1 = m.sub1 + 12

}

func main() {
	var name string
	var sub1, sub2, sub3, sub4 int
	var arr []Marks
	for i := 0; i < 3; i++ {
		fmt.Scanf("%v\n%v\n%v\n%v\n%v", &name, &sub1, &sub2, &sub3, &sub4)
		arr = append(arr, Marks{name, sub1, sub2, sub3, sub4})
		arr[i].change()
		fmt.Println(name, arr[i].average())
	}
	fmt.Println(arr)
}
