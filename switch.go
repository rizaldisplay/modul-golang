package main

func main() {
	name := "John Doe"

	switch name {
	case "John Doe":
		println("Hello, John!")
	case "Jane Doe":
		println("Hello, Jane!")
	default:
		println("Hello, Stranger!")
	}

	switch length := len(name); length > 5 {
	case true:
		println("Your name is quite long!")
	case false:
		println("Your name is short and sweet!")
	}

	name = "Alice"

	length := len(name)

	switch {
	case length > 5:
		println("Your name is quite long!")
	case length <= 5:
		println("Your name is short and sweet!")
	default:
		println("Unexpected length!")
	}
}
