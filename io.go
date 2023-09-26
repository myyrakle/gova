package java

import (
	"fmt"
	"os"
)

type PrintStream struct {
}

func (p PrintStream) Print(value any) {
	fmt.Print(value)
}

func (p PrintStream) Println(value any) {
	fmt.Println(value)
}

func (p PrintStream) Flush() {
	fmt.Println()
}

func (p PrintStream) Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

type InputStream struct {
}

// int read() => Reads the next byte of data from the input stream.
//
// int read(byte[] b) => Reads some number of bytes from the input stream and stores them into the buffer array b.
func (i InputStream) Read(data ...[]byte) int {

	switch len(data) {
	case 0:
		buffer := make([]byte, 1)

		if _, err := os.Stdin.Read(buffer); err != nil {
			return -1
		} else {
			return int(buffer[0])
		}
	case 1:
		if data, err := os.Stdin.Read(data[0]); err != nil {
			return -1
		} else {
			return data
		}
	default:
		panic("Too many arguments")
	}
}
