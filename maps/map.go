package main

import "fmt"

func main() {
	var m1 map[string]string // adding values Without use of make function.. panic: assignment to entry in nil map
	var key, value string
	m1 = make(map[string]string) // syntax : map[key_type]value_type

	m1 = map[string]string{"Id": "1", "Programming": "Go"}

	fmt.Println(m1)

	for i := 0; i < 3; i++ {
		fmt.Print("Enter the Key :")
		fmt.Scanln(&key)
		fmt.Print("Enter the Value :")
		fmt.Scanln(&value)
		m1[key] = value
		fmt.Println()
	}

	UpdateMap(m1)

	for k, v := range m1 {
		fmt.Printf("%v : %v\n", k, v)
	}

	//Delete a item from map
	fmt.Print("Enter Key to delete :")
	fmt.Scanln(&key)
	delete(m1, key)
	fmt.Println("Map after Deletion :", m1)

	//Check if Key is Available in Map
	fmt.Print("enter a key to check if Key Available :")
	fmt.Scanln(&key)
	if val, ok := m1[key]; ok {
		fmt.Println(key, "is available with value of", val)
	} else {
		fmt.Println(key, "is not available")
	}

}

func UpdateMap(m1 map[string]string) {
	m1["Location"] = "India"
}
