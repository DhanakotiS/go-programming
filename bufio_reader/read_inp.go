package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	/* func (b *Reader) ReadString(delim byte) (string, error) - Reads the given string*/
	str, err := reader.ReadString('\n')
	fmt.Print(str, err)
	/* >>> hello world
	-> hello world */

	/* func (b *Reader) ReadRune() (r rune, size int, err error) - Reads the byte value of first char called rune in int32 */
	runee, size, err := reader.ReadRune()
	fmt.Println(runee, size, err)
	/* >>> hello world
	->104 1 <nil> */

	/* func (b *Reader) ReadByte() (byte, error) - Read the byte value of first char*/
	byt, err := reader.ReadByte()
	fmt.Println(byt, err)
	/* >>> hello world
	->104 <nil> */

	/* func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error) - Reads byte values od provided string*/
	line, isPrefix, err := reader.ReadLine()
	fmt.Println(line, isPrefix, err)
	/* >>> hello world
	-> [104 101 108 108 111 32 119 111 114 108 100] false <nil> */

	/* func (b *Reader) ReadSlice(delim byte) (line []byte, err error) - Read the byte values of string */
	lin, err := reader.ReadSlice('\n')
	fmt.Println(lin, err)
	/* >>> hello world
	->[104 101 108 108 111 32 119 111 114 108 100 10] <nil>*/

	/* func (b *Reader) ReadBytes(delim byte) ([]byte, error) -Read the byte value of string */
	bytes, err := reader.ReadBytes('\n')
	fmt.Println(bytes, err)
	/* >>> hello world
	-> [104 101 108 108 111 32 119 111 114 108 100 10] <nil> */

	/* func (b *Reader) Size() int - return the size of reader buffer */
	n := reader.Size()
	fmt.Println(n)
	/* -> 4096 */

}
