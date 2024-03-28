package firstclassfunction

import "fmt"

func FunctionIIFE() {
	totalStudent := 30
	sudentPresent := 25

	var newPresent = func(total int, present int) float32 {
		result := (float32(present) / float32(total)) * 100

		return result

	}(totalStudent, sudentPresent) // fungsinya langsung di-invoke saat deklarasi

	fmt.Println("Attendance percentage :", newPresent)

	// Output:
	// Attendance percentage : 83.33333
}
