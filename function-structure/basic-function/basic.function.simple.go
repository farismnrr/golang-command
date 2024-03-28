package basicFunction

import "fmt"

// fungsi addNumbers
func addNumbersSimple() {
	n1 := 12
	n2 := 8

	sum := n1 + n2
	fmt.Println("Sum 12 + 8:", sum)
}

func FunctionSimple() {
	// memanggil fungsi addNumbers
	addNumbersSimple()
}
