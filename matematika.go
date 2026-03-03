package main

import "fmt"

func main() {
	var a = 10
	var b = 3
	var d = 5
	var c = a + b - d

	fmt.Println("Hasil dari a + b - d adalah:", c)

	var i = 10
	i += 10
	i += 5

	fmt.Println("Hasil dari i setelah penambahan adalah:", i)

	var j = 20
	j++
	fmt.Println("Hasil dari j setelah increment adalah:", j)

	var k = 15
	k--
	fmt.Println("Hasil dari k setelah decrement adalah:", k)
}
