package firstclassfunction

import "fmt"

func FunctionAsParameter() {
	// fungsi sebagai variable bernama `inc`
	inc := func(number int) int {
		return number + 1
	}
	// fungsi sebagai variable bernama `dec`
	dec := func(number int) int {
		return number - 1
	}

	fmt.Println(apply(inc, 5)) // 6
	fmt.Println(apply(dec, 5)) // 4
}

// fungsi apply membutuhkan 2 parameter, yaitu `anonymous function` dan `int`
func apply(f func(x int) int, number int) int {
	return f(number)
}
