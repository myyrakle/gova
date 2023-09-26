package java

import "fmt"

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
