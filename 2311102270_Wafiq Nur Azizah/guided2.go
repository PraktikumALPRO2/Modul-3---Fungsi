package main

import "fmt"

func main() {
	var sisi int
	sisi = inputSisi()

	luas := hitungLuas(sisi)
	keliling := hitungKeliling(sisi)

	tampilkanHasil(luas, keliling)
}

func inputSisi() int {
	var s int
	fmt.Print("Panjang sisi persegi: ")
	fmt.Scan(&s)
	return s
}

func hitungLuas(s int) int {
	return s * s
}

func hitungKeliling(s int) int {
	return 4 * s
}

func tampilkanHasil(luas, keliling int) {
	fmt.Println("Hasil luas persegi:", luas)
	fmt.Println("Hasil keliling persegi:", keliling)
}