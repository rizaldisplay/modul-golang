package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Surabaya"
	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Bandung", "Jawa Barat", "Indonesia"}
	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
