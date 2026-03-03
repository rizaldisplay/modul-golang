package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println("Nilai Akhir:", nilaiAkhir)
	fmt.Println("Absensi:", absensi)
	fmt.Println("Lulus Nilai Akhir:", lulusNilaiAkhir)
	fmt.Println("Lulus Absensi:", lulusAbsensi)
	fmt.Println("Lulus:", lulus)
}
