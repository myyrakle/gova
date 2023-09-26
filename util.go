package java

import (
	"fmt"
	"strconv"
)

type javaUtil struct {
	New_Scanner func(source InputStream) Scanner
}

var Util = javaUtil{
	New_Scanner: new_Scanner,
}

type Scanner struct {
	source InputStream
}

// TODO: File 인터페이스 추가해서 추상화하기
func new_Scanner(source InputStream) Scanner {
	return Scanner{
		source: source,
	}
}

func (this *Scanner) Next() string {
	buffer := []byte{}

	for {
		data := this.source.Read()

		if data == -1 {
			break
		}

		if data == '\n' {
			break
		}

		buffer = append(buffer, byte(data))
	}

	return string(buffer)
}

func (this *Scanner) NextInt() int {
	rawData := this.Next()

	if data, err := strconv.Atoi(rawData); err != nil {
		panic(&InputMismatchException{message: rawData})
	} else {
		return data
	}
}

type InputMismatchException struct {
	message string
}

func (this *InputMismatchException) Error() string {
	return fmt.Sprintf("InputMismatchException: %s", this.message)
}
