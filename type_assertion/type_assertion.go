package main

import "fmt"

type student struct {
	name string
}

func main() {

	var s1 interface{} = "Hello"
	val, isString := s1.(string)
	fmt.Println(val, isString)

	str, ok := s1.(float64) //no panic occurs
	fmt.Println(str, ok)

	// val = s1.(int) //panic
	// fmt.Println(val,isInt)

	st := student{"John"}

	var s interface{} = st

	struc, ok := s.(student)

	fmt.Println(struc, ok)
}
