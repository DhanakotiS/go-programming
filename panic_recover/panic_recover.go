package main

import "fmt"

// func recovering() {
// 		if r := recover(); r != nil {
// 			fmt.Println(r)
//			fmt.Println("Recovered from Panic!")
// 		}
// }

func main() {

	//defer recovering()
	defer func() {
		err_msg := recover()
		fmt.Println(err_msg)
		fmt.Println("Recovered from Panic!")
	}()

	var name string
	fmt.Scanln(&name)
	if len(name) <= 0 {
		panic("Panic : Name cannot be Empty!")
	}
	fmt.Println("Welcome", name)
}

/*
Panic : Name cannot be Empty!
Recovered from Panic!
*/
