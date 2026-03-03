package main

import "fmt"

func main() {
	// const firstName = "Afrizal"
	// const lastName = "Aminulloh"

	const (
		firtsname string = "Afrizal"
		lastname  string = "Aminulloh"
	)

	// error: cannot assign to firstName (constant "Afrizal" of type untyped string)
	// firstName = "Gojo"
	// lastName = "Satoru"

	fmt.Println("Hello, " + firtsname + " " + lastname + "!")
}
