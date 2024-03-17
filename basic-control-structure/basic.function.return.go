package basicControl

import "fmt"

// function definition
func addNumbersReturn(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func ControlFuncReturn() {
	// function call
	result := addNumbersReturn(21, 13)

	fmt.Println("Sum:", result)
}
