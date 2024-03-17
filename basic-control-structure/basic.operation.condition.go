package basicControl

import "fmt"

func ControlOperatorConditionComparison() {
	/*
		Operator	Keterangan
		>			Lebih dari
		<			Kurang dari
		>=			Lebih besar atau sama dengan
		<=			Lebih kecil atau sama dengan
		==			Sama dengan
		!=			Tidak sama dengan
	*/

	var name1, name2 = "John", "John"

	var result = name1 == name2
	var result2 = 10 < 2

	fmt.Println(result)
	fmt.Println(result2)

	// Output:
	/*
		true
		false
	*/
}

func ControlOperatorConditionLogic() {
	/*
		Operator	Keterangan
		&&			And / Dan
		\|\|		Or / Atau
		!			Not / Bukan / Kebalikan
	*/

	// Berikut hasil operasi dengan menggunakan Operator && :
	/*
		Nilai 1	Nilai 2	Hasil
		true	true	true
		true	false	false
		false	true	false
		false	false	false
	*/

	// Berikut hasil operasi dengan menggunakan Operator || :
	/*
		Nilai 1	Nilai 2	Hasil
		true	true	true
		true	false	true
		false	true	true
		false	false	false
	*/

	// Khusus untuk operasi !, kita bisa gunakan untuk membalikkan nilai boolean.
	/*
		Nilai	Hasil
		true	false
		false	true
	*/

	var score, attendance = 90, 70
	var graduated = score > 80 && attendance > 80

	fmt.Println(graduated)

	// Output: true
}

func ControlOperatorConditionIf() {
	var name = "Budi"

	if name == "Budi" {
		fmt.Println("Halo, Budi")
	}

	// Output:
	// Halo, Budi
}

func ControlOperatorConditionIfShort() {
	if x := 75; x > 70 {
		fmt.Println("Selamat, anda lulus")
	}

	// Output:
	// Selamat, anda lulus
}

func ControlOperatorConditionIfShort2() {
	if x, y := 75, 80; x > 70 && y > 70 {
		fmt.Println("Selamat, anda lulus")

		fmt.Println("Nilai rata-rata:", (x+y)/2)
	}

	// Output:
	/*
		Selamat, anda lulus
		Nilai rata-rata: 77
	*/
}

func ControlOperatorConditionElse() {
	var name = "Anonim"

	if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Helo, boleh kenalan ?")
	}

	// Output:
	// Helo, boleh kenalan ?
}

func ControlOperatorConditionElseIf() {
	var score = 81

	if score == 100 {
		fmt.Println("Perfect")
	} else if score >= 90 && score < 100 {
		fmt.Println("Fantastic")
	} else if score >= 80 && score < 90 {
		fmt.Println("Good")
	} else if score >= 70 && score < 80 {
		fmt.Println("Graduated")
	} else {
		fmt.Println("Failed")
	}

	// Output: Good
}

func ControlOperatorConditionNestedCondition() {
	num := 1

	if num%2 == 0 {
		if num%3 == 0 {
			fmt.Println("Bilangan genap dan kelipatan 3")
		} else {
			fmt.Println("Bilangan genap bukan kelipatan 3")
		}
	} else {
		fmt.Println("Bilangan ganjil")
	}
	//Output: Bilangan Ganjil
}

func ControlOperatorConditionSwitchCase() {
	day := 4
	switch day {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Invalid day")
	}
	// Output: Kamis
}

func ControlOperatorConditionSwitchCase2() {
	day := 4
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
	// Output: Weekday
}

func ControlOperatorConditionInCase() {
	number := 6
	switch {
	case number > 10:
		fmt.Println("High number")
	case number > 5:
		fmt.Println("Middle number")
	case number > 1:
		fmt.Println("Low number")
	default:
		fmt.Println("Zero")
	}
	// Output: Middle number
}

func ControlOperatorConditionFallthrough() {
	number := 6
	switch {
	case number > 10:
		fmt.Println("High number")
		fallthrough // check to the next case
	case number > 5:
		fmt.Println("Middle number")
		fallthrough // check to the next case
	case number > 1:
		fmt.Println("Low number")
		fallthrough // check to the next case
	default:
		fmt.Println("Zero")
	}
	// Output:
	/*
		Middle number
		Low number
		Zero
	*/
}
