package main

func getFullName(firstName string, lastName string) (string, string) {
	fullName := firstName + " " + lastName
	reversedFullName := lastName + " " + firstName
	return fullName, reversedFullName
}

func main() {
	fullName, reversedFullName := getFullName("Alice", "Smith")
	println("Full Name:", fullName)
	println("Reversed Full Name:", reversedFullName)
}
