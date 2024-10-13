package main

import "fmt"

func main() {
	var sisi int
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&sisi)

	luas := hitungLuas(sisi)
	keliling := hitungKeliling(sisi)

	fmt.Println("Luas persegi:", luas)
	fmt.Println("Keliling persegi:", keliling)
}

func hitungLuas(sisi int) int {
	return sisi * sisi
}

func hitungKeliling(sisi int) int {
	return 4 * sisi
}