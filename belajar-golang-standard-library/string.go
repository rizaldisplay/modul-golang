package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Afrizal Aminulloh", "Afrizal"))
	fmt.Println(strings.Split("Afrizal Aminulloh", " "))
	fmt.Println(strings.ToLower("Afrizal Aminulloh"))
	fmt.Println(strings.ToUpper("Afrizal Aminulloh"))
	fmt.Println(strings.Trim("      Afrizal Aminulloh", " "))
	fmt.Println(strings.ReplaceAll("Afrizal Aminulloh      ", " Aminulloh ", " RUDI"))
}
