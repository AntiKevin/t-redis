package resp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Deserializer is a function that reads a RESP message from a buffer and parses it into a string.
func Deserializer() {

	input := "skdhngkjasfdh"
	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()

	if b != '$' {
		fmt.Println("Invalid type, expecting bulk strings only")
		os.Exit(1)
	}

	size, _ := reader.ReadByte()

	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	// consume /r/n
	reader.ReadByte()
	reader.ReadByte()

	name := make([]byte, strSize)
	reader.Read(name)

	println("Name: ", string(name))
}
