package main

import "fmt"

func arg_input_scan() {
	var name, address string
	fmt.Print("Enter your name and address : ")
	fmt.Scan(&name, &address) // set input to variable 'name' and 'address'

	fmt.Println("Hello", name)
	fmt.Println("Your address", address)

	// Output:
	/*
		Enter your name and address : Budi Jakarta
		Hello Budi
		Your address Jakarta
	*/
}

func arg_input_scanln() {
	var name, address string
	fmt.Print("Enter your name : ")
	fmt.Scanln(&name)

	fmt.Print("Enter your address : ")
	fmt.Scanln(&address)

	fmt.Println("Hello", name)
	fmt.Println("Your address", address)

	// Output:
	/*
		Enter your name : Budi
		Enter your address : Jakarta
		Hello Budi
		Your address Jakarta
	*/
}

func arg_input_scanf() {
	var name, address string

	fmt.Print("Enter your name and address : ")
	fmt.Scanf("%s %s", &name, &address) // menerima 2 input string yang dipisahkan oleh spasi, yang pertama akan diassign ke variable name dan yang kedua akan diassign ke variable address

	fmt.Println("Hello", name)
	fmt.Println("Your address", address)

	// Output:
	/*
		Enter your name and address : Budi Jakarta
		Hello Budi
		Your address Jakarta
	*/
}
