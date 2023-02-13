package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAKhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80

	var lulus bool = lulusAbsensi && lulusNilaiAKhir

	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 90 && absensi >= 80)
}
