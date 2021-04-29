package main

import "fmt"

func main() {
	x := make(map[string]string)
	x["H"] = "Hydrogen"
	x["He"] = "Helium"
	x["Li"] = "Lithium"
	x["B"] = "Boron"
	x["Be"] = "Beryllium"

	fmt.Println(x["Li"])
	fmt.Println("Before deletion", x["B"])
	delete(x, "B")
	fmt.Println("After deletion", x["B"])

}
