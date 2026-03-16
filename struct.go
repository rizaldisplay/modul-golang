package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var rizal Customer
	fmt.Println(rizal)

	rizal.Name = "Afrizal Aminulloh"
	rizal.Address = "Indonesia"
	rizal.Age = 30
	fmt.Println(rizal)
	fmt.Println(rizal.Name)
	fmt.Println(rizal.Address)
	fmt.Println(rizal.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	joko.sayHello("Agus")
}
