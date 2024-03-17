package basicCommand

import "fmt"

func CommandDataTypeStr() {
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

func CommandDataTypeStrFuncLen() {
	fmt.Println(len("Hello World"))
	// Output:
	// 11 # total character
}

func CommandDataTypeStrFuncGetChar() {
	fmt.Println("Golang"[0])
	fmt.Println("Golang"[2])
	// Output:
	// 71 # output is ASCII code of "G"
	// 108 # output is ASCII code of "l"

	fmt.Println(string("Golang"[0]))
	// Output:
	// G
}
