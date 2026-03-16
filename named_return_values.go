package main

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Alice"
	middleName = "B."
	lastName = "Smith"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	println("First Name:", firstName)
	println("Middle Name:", middleName)
	println("Last Name:", lastName)
}
