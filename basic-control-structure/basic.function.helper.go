package basicControl

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// panggil fungsi reverse sebagai fungsi helper
	xReversed := reverse(x)

	return xReversed == x
}

func reverse(x int) int {
	var reversedDigit int

	for x != 0 {
		lastDigit := x % 10
		reversedDigit = reversedDigit*10 + lastDigit

		x = x / 10
	}

	return reversedDigit
}

func ControlFuncHelper() {

	output := isPalindrome(121)
	fmt.Println(output)

	output = isPalindrome(12)
	fmt.Println(output)

	output = isPalindrome(-101)
	fmt.Println(output)
}
