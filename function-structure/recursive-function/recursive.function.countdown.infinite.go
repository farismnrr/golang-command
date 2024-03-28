package recursivefunction

import "fmt"

func countDownRecursive1(number int) {
	// menampilkan parameter `number`
	fmt.Println(number)
	// pemanggilan recursive dengan mengurangi nilai `number`
	countDownRecursive1(number - 1)
}

func FunctionRecursiveCountdownInfinite() {
	fmt.Println("Countdown Starts:")
	// pemanggilan fungsi `countDown()`
	countDownRecursive1(3)

	// Output:
	/*
		Countdown Starts:
		3
		2
		1
		0
		-1
		-2
		-3
		-4
		...
	*/
}
