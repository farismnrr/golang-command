package basicControl

import (
	"fmt"
	"strings"
)

func myHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func ControlFuncParameterSimpleVariadic() {
	myHobbies("Aditira", "Coding", "Gaming", "Jogging")
}
