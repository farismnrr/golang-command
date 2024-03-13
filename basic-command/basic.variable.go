package basicCommand

import "fmt"

func VarBasicDeclaration() {
	// Method 1
	var hello string
	var age int
	var pi float64

	hello = "Hello Go!" // assign new value
	age = 15
	pi = 3.14

	fmt.Println(hello, age, pi)

	// Method 2
	var hello2 = "Hello Go!"
	var hello3 = "Hello Go!"
	// sama saja dengan
	// var hello string = "Hello Go!"

	fmt.Println(hello2, hello3)
}

func VarDefaultValue() {
	/*
		Tipe Data		Default Value
		string			""
		int			0
		bool			false
	*/

	var (
		name      string
		age       int
		isMarried bool
	)

	fmt.Println(name)      // ""
	fmt.Println(age)       // 0
	fmt.Println(isMarried) // false
}

func VarShortDeclaration() {
	hello := "Hello World"
	// sama saja dengan
	// var hello string = "Hello World"

	fmt.Println(hello)
}

func ConstDeclaration() {
	const pi float64 = 3.14159265359
	// jika kita melakukan assign nilai baru, maka error
	// pi = 2.34 --> error
}

func VarConstMultiDeclaration() {
	// Method 1
	var name string
	var address string
	var status string

	// can be simplified to
	var name2, address2, status2 string

	// Method 2
	var (
		name3      string
		age3       int
		isMarried3 bool
	)

	const (
		Pi float64 = 3.14
		E  float64 = 2.718
	)

	// Method 3
	name4, age4, isMarried4 := "Budi", 30, true

	fmt.Println(name, name2, name3, name4, age3, age4, address, address2, status, status2, isMarried3, isMarried4)
}

func VarUnderscore() {
	hello, _ := "Hello", "Go!"

	_, _ = fmt.Scanf("%s", &hello)
}

func VarNamingRule() {
	var myVariableName = "Hello Go!"   // camel case
	var MyVariableName = "Hello Go!"   // pascal case
	var my_variable_name = "Hello Go!" // snake case

	fmt.Println(myVariableName, MyVariableName, my_variable_name)
}
