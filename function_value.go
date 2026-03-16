package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye, " + name + "!"
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("Alice"))
	fmt.Println(misal("Bob"))

}
