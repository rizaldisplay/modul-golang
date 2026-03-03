package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "John"
	name[1] = "Doe"
	name[2] = "Smith"

	fmt.Println("Nama 1:", name[0])
	fmt.Println("Nama 2:", name[1])
	fmt.Println("Nama 3:", name[2])

	// make in dirrect
	name2 := [3]string{"Alice", "Bob", "Charlie"}

	fmt.Println(name2)
}
