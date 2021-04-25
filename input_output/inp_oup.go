package main

import "fmt"

func main() {
	fmt.Println("Enter First Name:")
	var fname string
	fmt.Scanln(&fname)

	fmt.Print("Enter the Last Name:")
	var lname string
	fmt.Scan(&lname)

	fmt.Printf("%s", "Enter the Age:")
	var age int
	fmt.Scanf("%d", &age)

	fmt.Println(fname, lname, age)
	fmt.Print(fname, lname, age, "\n")
	fmt.Printf("%s %s %d \n", fname, lname, age)
}

// Enter First Name:
// John
// Enter the Last Name:Doe
// Enter the Age:26
// John Doe 26
// JohnDoe26
// John Doe 26
