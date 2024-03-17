package basicCommand

import "fmt"

func CommandArgumentCommentInline() {
	// this is comment
	// print output "Hello World!"
	// fmt.Println("Hello World!");

	fmt.Println("Hello Go!")
}

func CommandArgumentCommentMultiline() {
	/*
		this is comment
		print output "Hello World!"
		fmt.Println("Hello World!");
	*/

	fmt.Println("Hello Go!")
}
