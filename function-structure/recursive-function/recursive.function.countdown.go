package recursivefunction

import "fmt"

/*
	Number > 0	Print	Recursive Call
	true	3	countDown(2)
	true	2	countDown(1)
	true	1	countDown(0)
	false	Countdown stops	function execution stops
*/

func countDownRecursive2(number int) {
	if number > 0 {
		fmt.Println(number)
		countDownRecursive2(number - 1)
	} else {
		fmt.Println("Countdown Stops")
	}
}

func FunctionRecursiveCountdown() {
	fmt.Println("Countdown Starts")
	countDownRecursive2(3)

	// Output:
	/*
		Countdown Starts
		3
		2
		1
		Countdown Stops
	*/
}
