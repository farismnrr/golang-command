package recursivefunction

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func RecursiveFactorial() {
	var number = 3
	var result = factorial(number)
	fmt.Println("The factorial of 3 is", result)

	// Output: The factorial of 3 is 6
}
