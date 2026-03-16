package main

func sayHello(name string, greeting string) {
	println(greeting + ", " + name + "!")
}

func main() {
	sayHello("Alice", "Hi")
	sayHello("Bob", "Hey")
}
