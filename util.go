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

func (this *Scanner) NextByte() Byte {
	rawData := this.Next()

	if data, err := strconv.Atoi(rawData); err != nil {
		panic(&InputMismatchException{message: rawData})
	} else {
		return Byte(data)
	}
}

func (this *Scanner) NextShort() Short {
	rawData := this.Next()

	if data, err := strconv.Atoi(rawData); err != nil {
		panic(&InputMismatchException{message: rawData})
	} else {
		return Short(data)
	}
}

func (this *Scanner) NextInt() Integer {
	rawData := this.Next()

	if data, err := strconv.Atoi(rawData); err != nil {
		panic(&InputMismatchException{message: rawData})
	} else {
		return Integer(data)
	}
}

func (this *Scanner) NextLong() Long {
	rawData := this.Next()

	if data, err := strconv.Atoi(rawData); err != nil {
		panic(&InputMismatchException{message: rawData})
	} else {
		return Long(data)
	}
}

func (this *Scanner) NextFloat() Float {
	rawData := this.Next()

	if data, err := strconv.ParseFloat(rawData, 32); err != nil {
		panic(&InputMismatchException{message: rawData})
	} else {
		return Float(data)
	}
}

func (this *Scanner) NextDouble() Double {
	rawData := this.Next()

	if data, err := strconv.ParseFloat(rawData, 64); err != nil {
		panic(&InputMismatchException{message: rawData})
	} else {
		return Double(data)
	}
}

func (this *Scanner) NextBoolean() Boolean {
	rawData := this.Next()

	if data, err := strconv.ParseBool(rawData); err != nil {
		panic(&InputMismatchException{message: rawData})
	} else {
		return Boolean(data)
	}
}

type InputMismatchException struct {
	message string
}

func (this *InputMismatchException) Error() string {
	return fmt.Sprintf("InputMismatchException: %s", this.message)
}
