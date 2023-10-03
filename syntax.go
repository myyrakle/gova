package java

type tryCatch struct {
	err any
}

func Throw(exception any) {
	panic(exception)
}

func Try(callback func()) (t tryCatch) {
	defer func() {
		t.err = recover()
	}()
	callback()

	return
}

func (t tryCatch) Catch(callback func(exception any)) {
	if t.err != nil {
		callback(t.err)
	}
}
