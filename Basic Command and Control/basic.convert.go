package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func conversion_between_numbering_data_type() {
	var number int32 = 10
	var bigNum = int64(number)     // convert int32 to int64
	var floatNum = float32(number) // convert int32 to float32

	fmt.Println(number, bigNum, floatNum)
}

func conversion_to_string() {
	var number int32 = 10
	var isMaried bool = false

	var numString = fmt.Sprintf("%d", number)        // convert int32 ke string
	var isMariedString = fmt.Sprintf("%t", isMaried) // convert bool ke string
	var Pi = fmt.Sprintf("%f", math.Pi)              // convert float64 dari variable 'Pi' di package 'math' ke string

	fmt.Println(numString)
	fmt.Println(isMariedString)
	fmt.Println(Pi)

	fmt.Println("type numString:", reflect.TypeOf(numString))
	fmt.Println("type isMarriedString:", reflect.TypeOf(isMariedString))
	fmt.Println("type Pi:", reflect.TypeOf(Pi))

	// Output:
	/*
		10
		false
		3.141592653589793
		type numString: string
		type isMarriedString: string
		type Pi: string
	*/
}

func conversion_int_to_string() {
	var str = strconv.Itoa(10) // convert string to int

	fmt.Println(str)
	fmt.Println("type:", reflect.TypeOf(str))

	// Output:
	/*
		10
		type: string
	*/
}

func conversion_string_to_int() {
	var num, _ = strconv.Atoi("10") // convert string to int

	fmt.Println(num)
	fmt.Println("type:", reflect.TypeOf(num))

	// Output:
	/*
		10
		type: int
	*/
}
