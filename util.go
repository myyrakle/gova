package java

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
