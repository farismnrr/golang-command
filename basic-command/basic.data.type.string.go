package basicCommand

import "fmt"

func DataTypeString() {
	// "Hello World"
	// "2134 !_+@&#*()"

	/*
		`
		This is string
		With multi-line
		Hello Go!
		`
	*/
}

func DataTypeFuncLen() {
	fmt.Println(len("Hello World"))
	// Output:
	// 11 # total character
}

func DataTypeFuncGetChar() {
	fmt.Println("Golang"[0])
	fmt.Println("Golang"[2])
	// Output:
	// 71 # output is ASCII code of "G"
	// 108 # output is ASCII code of "l"

	fmt.Println(string("Golang"[0]))
	// Output:
	// G
}
