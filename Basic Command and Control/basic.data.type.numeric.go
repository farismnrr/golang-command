package main

import "fmt"

func numeric_int() {
	/* Tipe data	Nilai minimum	Nilai maksimum
	int8	-128	127
	int16	-32768	32767
	int32	-2147483648	2147483647
	int64	-9223372036854775808	9223372036854775807 */

	fmt.Println(-30)
	fmt.Println(10_000_000_000)
	// Output:
	// -30
	// 10000000000
}

func numeric_uint() {
	/* Tipe data	Nilai minimum	Nilai maksimum
	uint8	0	255
	uint16	0	65535
	uint32	0	4294967295
	uint64	0	18446744073709551615 */
}

func numeric_float() {
	/* Tipe data	Nilai minimum	Nilai maksimum
	float32	-3.4E+38	3.4E+38
	float64	-1.7E+308	1.7E+308 */

	fmt.Println(3.14)
	fmt.Println(-2.232321431412)
	// Output:
	// 3.14
	// -2.232321431412
}

func numeric_operation() {
	/* Operator / Syntax	Keterangan
	+	Penambahan
	-	Pengurangan
	*	Perkalian
	/	Pembagian
	%	Mendapatkan sisa pembagian */

	fmt.Println(1 + 1)
	fmt.Println(2 - 1)
	fmt.Println(2 * 2)
	fmt.Println(4 / 2)
	fmt.Println(10 % 3) // sisa pembagian 10 dengan 3 adalah 1
	// Output:
	/*
		2
		1
		4
		2
		1
	*/
}

func numeric_augmented_assignment() {
	/* Operator / Syntax	Operasi matematika	Keterangan
	a += 10	a = a + 10	Penambahan
	a -= 10	a = a - 10	Pengurangan
	a *= 10	a = a * 10	Perkalian
	a /= 10	a = a / 10	Pembagian
	a %= 10	a = a % 10	Mendapatkan sisa pembagian */

	var a = 10
	a += 10 // a = 10 + 10 = 20
	fmt.Println(a)
	// Output:
	// 20
}

func numeric_unary_operator() {
	/* Operator / Syntax	Operasi matematika	Keterangan
	a++	a = a + 1	increment (naik 1 angka)
	a--	a = a - 1	decrement (turun 1 angka)
	-a	a = -a	mengubah jadi negatif (negasi)
	+a	a = +a	mengubah jadi positif */

	var a = 10
	a++ // a = 10 + 1 = 11

	fmt.Println(a)

	var b = 10
	fmt.Println(-b)
	// Output:
	// 11
	// -10
}
