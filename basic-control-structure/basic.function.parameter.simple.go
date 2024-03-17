package basicControl

import "fmt"

// fungsi dengan satu parameter
func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func ControlFuncParameterSimple() {
	// mengirim parameter yang berbeda dalam fungsi yang sama
	sayHello("Aditira")
	sayHello("Dito")
	sayHello("Dina")
}
