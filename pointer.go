package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// address2 := &address1

	// address2.City = "Surabaya"

	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah menjadi Surabaya

	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Surabaya"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi surabaya
}
