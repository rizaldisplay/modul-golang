package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, error := strconv.ParseBool("true")
	if error != nil {
		fmt.Println("error:", error.Error())
	} else {
		fmt.Println("result:", result)
	}

	resultInt, error := strconv.Atoi("1000")
	if error != nil {
		fmt.Println("error:", error.Error())
	} else {
		fmt.Println("result:", resultInt)
	}

	binary := strconv.FormatInt(100, 2)
	fmt.Println("binary:", binary)

	var stringInt string = strconv.Itoa(100)
	fmt.Println("stringInt:", stringInt)
}
