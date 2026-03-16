package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip the rest of the loop body for even numbers
		}

		fmt.Println("Current value of i:", i)
	}
}
