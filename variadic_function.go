package main

func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result1 := sumAll(1, 2, 3)
	println("Sum of 1, 2, 3:", result1)

	result2 := sumAll(4, 5, 6, 7)
	println("Sum of 4, 5, 6, 7:", result2)
}
