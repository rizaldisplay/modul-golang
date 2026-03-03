package main

import "fmt"

func main() {
	// var name string

	name := "Afrizal Gojo"
	fmt.Println("Hello, " + name + "!")

	name = "Afrizal Aminulloh"
	fmt.Println("Hello, " + name + "!")

	var (
		firstName = "Afrizal"
		lastName  = "Aminulloh"
	)

	fmt.Println("Hello, " + firstName + " " + lastName + "!")
}
