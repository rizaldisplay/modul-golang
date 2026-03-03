package main

import "fmt"

func main() {
	names := [...]string{"John", "Doe", "Smith", "Alice", "Bob", "Charlie"}

	slice1 := names[0:3] // Slice dari index 0 hingga 2
	fmt.Println("Slice 1:", slice1)

	slice2 := names[3:6] // Slice dari index 3 hingga 5
	fmt.Println("Slice 2:", slice2)

	slice3 := names[1:5] // Slice dari index 1 hingga 4
	fmt.Println("Slice 3:", slice3)

	var slice4 []string = names[:] // Slice seluruh array
	fmt.Println("Slice 4:", slice4)

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	daySlice1 := days[5:]
	fmt.Println("Day Slice 1:", daySlice1)

	daySlice1[0] = "Weekend" // Mengubah nilai pada slice akan mempengaruhi array asli
	daySlice1[1] = "Holiday"

	fmt.Println("Day Slice 1 setelah diubah:", daySlice1)
	fmt.Println("Array Days setelah diubah:", days)

	dayslice2 := append(daySlice1, "New Holiday") // Menambahkan elemen ke slice, tidak mempengaruhi array asli
	fmt.Println("Day Slice 2:", dayslice2)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Hello"
	newSlice[1] = "World"

	fmt.Println("New Slice:", newSlice)
	fmt.Println(len(newSlice)) // Panjang slice
	fmt.Println(cap(newSlice)) // Kapasitas slice

	newSlice2 := append(newSlice, "Go", "Programming")
	fmt.Println("New Slice 2:", newSlice2)
	fmt.Println(len(newSlice2)) // Panjang slice setelah append
	fmt.Println(cap(newSlice2)) // Kapasitas slice setelah append

	newSlice2[0] = "Hi" // Mengubah nilai pada newSlice2 tidak mempengaruhi newSlice
	fmt.Println("New Slice 2 setelah diubah:", newSlice2)
	fmt.Println("New Slice setelah newSlice2 diubah:", newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println("From Slice:", fromSlice)
	fmt.Println("To Slice:", toSlice)
}
