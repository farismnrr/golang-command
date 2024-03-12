package main

import "fmt"

func println() {
	fmt.Println("This is output Println")
	fmt.Println("Hello Go!")
	fmt.Println()
	fmt.Println("Hello", 1234, true)

	// Output
	/*
		This is output Println
		Hello Go!

		Hello 1234 true
	*/
}

func print() {
	fmt.Print("This is output Print ")
	fmt.Print("Hello Go!")

	// Output:
	// This is output Print Hello Go!

	fmt.Print("This is output Print\n")
	fmt.Print("Hello Go!")
	// Output:
	/* This is output Print
	Hello Go! */
}

func printf() {
	fmt.Printf("My name is %s\n", "Budi")           // %s is formatting for string data type
	fmt.Printf("%d is my contact number\n", 123434) // %d is formatting for integer / numberic data type
	// Output:
	/* My name is Budi
	123434 is my contact number */
}
