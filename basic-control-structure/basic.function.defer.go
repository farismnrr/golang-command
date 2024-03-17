package basicControl

import "fmt"

func ControlFuncDefer() {
	fmt.Println(countByOne(2))
}

func countByOne(x int) int {
	defer fmt.Println("Counting done")

	fmt.Println("Counting start")

	return x + 1
}
