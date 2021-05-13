package main

import (
	"flag"
	"fmt"
)

func main() {
	var fInt = flag.Int("year", 2021, "Current year")
	flag.Parse()

	year := *fInt

	if year%4 == 0 && year%400 == 0 {
		fmt.Println("The given year", year, "is a Leap year")
	} else {
		fmt.Println("The given year", year, "is not a Leap year")
	}

}

/*
>>> go run flag.go --year=2100
The given year 2100 is not a Leap year
*/
