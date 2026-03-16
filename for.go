package main

import "fmt"

func main() {
	counter := 1

	// for counter <= 5 {
	// 	fmt.Println("Counter:", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Counter:", counter)
	}

	fmt.Println("Counter has reached", counter)

	names := []string{"Alice", "Bob", "Charlie"}

	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	for _, name := range names {
		fmt.Printf("Name: %s\n", name)
	}

}
