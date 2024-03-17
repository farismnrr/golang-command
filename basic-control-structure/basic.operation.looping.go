package basicControl

import (
	"fmt"
	"time"
)

func ControlOperatorLoopingBasic() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello Go!", i+1)
	}

	// Output:
	/*
		Hello Go! 1
		Hello Go! 2
		Hello Go! 3
		Hello Go! 4
		Hello Go! 5
	*/
}

func ControlOperatorLoopingBasic2() {
	var i int

	for i = 0; i < 5; i++ {
		fmt.Println("Hello Go!", i+1)
	}

	// Output:
	/*
		Hello Go! 1
		Hello Go! 2
		Hello Go! 3
		Hello Go! 4
		Hello Go! 5
	*/
}

func ControlOperatorLoopingBasic3() {
	for i, total := 1, 0; i <= 5; i++ {
		total += i // 1 + 2 + 3 + 4 + 5 = 15
		fmt.Println(i, "+", total-i, "=", total)
	}

	// Output:
	/*
		1 + 0 = 1
		2 + 1 = 3
		3 + 3 = 6
		4 + 6 = 10
		5 + 10 = 15
	*/
}

func ControlOperatorLoopingBreak() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Hello Go!", i+1)
	}

	// Output
	/*
		Hello Go! 1
		Hello Go! 2
		Hello Go! 3
	*/
}

func ControlOperatorLoopingContinue() {
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Hello Go!", i)
	}
	// Output:
	/*
		Hello Go! 1
		Hello Go! 3
		Hello Go! 5
	*/
}

func ControlOperatorLoopingConditionOnly() {
	i := 0 // init statement
	for i < 5 {
		fmt.Println("Hello Go!", i+1)
		i++ // post statement
	}

	// Output
	/*
		Hello Go! 1
		Hello Go! 2
		Hello Go! 3
		Hello Go! 4
		Hello Go! 5
	*/
}

func ControlOperatorLoopingEmptyForStatement() {
	var i = 0 // init statement

	for {
		if i >= 5 { // condition statement
			break // menghentikan looping jika kondisi 'true'
		}

		fmt.Println("Hello Go!", i+1) // code execute

		i++ // post statement
	}

	// Output:
	/*
		Hello Go! 1
		Hello Go! 2
		Hello Go! 3
		Hello Go! 4
		Hello Go! 5
	*/
}

func ControlOperatorLoopingInfiniteLoop() {
	for {
		fmt.Println("Hello Go!")

		time.Sleep(1 * time.Second) // ini adalah fungsi yang dibuat agar program delay selama 1 detik
	}
}

func ControlOperatorLoopingNestedLooping() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Hello Go! i: %d, j: %d\n", i, j)
		}
	}

	// Output:
	/*
		Hello Go! i: 1 j: 1
		Hello Go! i: 1 j: 2
		Hello Go! i: 1 j: 3
		Hello Go! i: 2 j: 1
		Hello Go! i: 2 j: 2
		Hello Go! i: 2 j: 3
	*/
}
