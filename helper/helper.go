package helper

import "fmt"

var version = "1.0.0"
var Application = "golang helper"

func SayGoodbye(name string) string {
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() { SayGoodbye("Rizal"); fmt.Println(version) }
