package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	b := []byte("\n")

	writer := bufio.NewWriter(os.Stdout)

	// 'Size' returns the size of buffer
	fmt.Println("Default Writer Size:", writer.Size())

	writer = bufio.NewWriterSize(os.Stdin, 8192)

	fmt.Println("Custom Writer Size:", writer.Size())

	writer.Write(b)

	writer.WriteByte(104)

	writer.WriteRune(112)

	writer.WriteString("Hello\n")

	// 'Buffered' returns no of bytes currently in buffer
	fmt.Println("Buffered Size:", writer.Buffered())
	
	// 'Available' returns remaining size in buffer
	fmt.Println("Available Size", writer.Available())

	// 'Flush' Writes the buffered data into underlying io.writer
	writer.Flush()

	fmt.Println("Reset Demo")

	writer.WriteByte(104)
	writer.Write(b)

	writer.WriteRune(112)
	writer.Write(b)

	fmt.Println("Buffered Size:", writer.Buffered())

	writer.Flush()

	writer.Reset(os.Stdout)

	fmt.Println("Buffered Size after Reset:", writer.Buffered())

}
