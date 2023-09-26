package java

type javaLang struct {
	System javaLangSystem
}

type javaLangSystem struct {
	Out PrintStream
	Err PrintStream
	In  InputStream
}

var Lang = javaLang{}
