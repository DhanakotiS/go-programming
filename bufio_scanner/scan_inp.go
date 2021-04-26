package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var buf []byte

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Buffer(buf, 4096)

	scanner.Scan()

	fmt.Println(scanner.Text())

	fmt.Println(scanner.Bytes())

	fmt.Println(scanner.Err())

}
