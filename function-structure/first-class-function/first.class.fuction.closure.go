package firstclassfunction

import "fmt"

func NewCounter() func() int {
	i := 0 // nilai variabel `i` adalah closure bagi inner function `inc`
	inc := func() int {
		i++
		return i
	}
	return inc
}

func FunctionClosure() {
	ctr := NewCounter()

	fmt.Println(ctr()) // 1 => nilai variabel `i` sekarang 0 dan menghasilkan 1
	fmt.Println(ctr()) // 2 => nilai variabel `i` sekarang 1 dan menghasilkan 2
}
