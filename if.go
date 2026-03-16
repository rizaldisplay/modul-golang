package main

import "fmt"

func main() {
	name := "John Doe"

	if name == "John Doe" {
		fmt.Println("Hello, John!")
	} else if name == "Jane Doe" {
		fmt.Println("Hello, Jane!")
	} else {
		fmt.Println("Hello, Stranger!")
	}

	length := len(name)

	if length > 10 {
		fmt.Println("Your name is quite long!")
	} else {
		fmt.Println("Your name is short and sweet!")
	}

}
