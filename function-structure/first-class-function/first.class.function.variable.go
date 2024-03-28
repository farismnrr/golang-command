package firstclassfunction

import "fmt"

func FunctionAsVariable() {
	// Anonymous function assigned to variable inc:
	inc := func(number int) int {
		return number + 1
	}

	// Assigned variable function to variable calculate:
	calculate := inc(5)
	fmt.Println(calculate) // 6
}
