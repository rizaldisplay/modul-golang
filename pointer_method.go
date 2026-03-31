package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rizal := &Man{Name: "Rizal"}
	rizal.Married()

	fmt.Println(rizal.Name)
}
