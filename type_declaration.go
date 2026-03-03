package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRizal NoKTP = "111111111111111"

	var contoh string = "22222222222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println("ktpRizal:", ktpRizal)
	fmt.Println("contohKTP:", contohKTP)
}
