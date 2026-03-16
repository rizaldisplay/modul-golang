package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name":    "John Doe",
		"age":     "30",
		"country": "USA",
	}

	fmt.Println("Person Map:", person)
	fmt.Println("Name:", person["name"])

	person["age"] = "31"
	fmt.Println("Updated Person Map:", person)

	fmt.Println("Number of Elements in Map:", person["null"])

	book := make(map[string]string)
	book["title"] = "The Great Gatsby"
	book["author"] = "F. Scott Fitzgerald"
	book["year"] = "1925"

	fmt.Println("Book Map:", book)
	delete(book, "year")
	fmt.Println("Book Map after deletion:", book)

}
