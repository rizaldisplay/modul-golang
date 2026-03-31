package main

import (
	"fmt"
	"project-golang/helper"
)

func main() {
	result := helper.SayHello("Rizal")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.version)             // tidak bisa diakses karena variable version tidak diexport (diawali dengan huruf kecil)
	fmt.Println(helper.SayGoodbye("Rizal")) // tidak bisa diakses karena function SayGoodbye tidak diexport (diawali dengan huruf kecil)
}
