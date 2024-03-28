package basicFunction

import "fmt"

// fungsi menggunakan 2 parameter
func addNumbersMulti(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("Sum:", sum)
}

func FunctionParameterMulti() {
	// mengirim 2 parameter pada fungsi addNumbers
	addNumbersMulti(21, 13)
}
