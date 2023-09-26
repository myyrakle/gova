# gova

![](https://img.shields.io/badge/language-Go-00ADD8) ![](https://img.shields.io/badge/version-v0.1.0-brightgreen) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

Go is actually Java.
  
[document](https://pkg.go.dev/github.com/myyrakle/gova)

## install

```
go get "github.com/myyrakle/gova@{version}"
```

## how to use

like this

```
package main

import java "github.com/myyrakle/gova"

func main() {
	java.Lang.System.Out.Println("Hello, World!")

	java.Lang.System.Out.Print("input integer number: ")
	scanner := java.Util.New_Scanner(java.Lang.System.In)

	input := scanner.NextInt()
	java.Lang.System.Out.Println(input)
}
```

![image](https://github.com/myyrakle/gova/assets/16988115/8dc660ad-1b1f-4e3b-ba71-c084e016614a)
