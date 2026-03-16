package main

import "fmt"

func sayHelloTo(firstName string, lastName string) string {
	return fmt.Sprintf("Hello, %s %s!", firstName, lastName)
}

func main() {
	greeting := sayHelloTo("Alice", "Smith")
	fmt.Println(greeting)
}
