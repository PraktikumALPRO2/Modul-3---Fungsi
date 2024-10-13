package main

import (
	"fmt"
)

// Fungsi untuk menghitung luas persegi panjang
func hitungLuas(panjang, lebar float64) float64 {
	return panjang * lebar
}

// Fungsi untuk menghitung keliling persegi panjang
func hitungKeliling(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func main() {
	var panjang, lebar float64

	// Input nilai panjang dan lebar
	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scanln(&lebar)

	// Menghitung luas dan keliling
	luas := hitungLuas(panjang, lebar)
	keliling := hitungKeliling(panjang, lebar)

	// Output hasil perhitungan
	fmt.Printf("Luas persegi panjang: %.2f\n", luas)
	fmt.Printf("Keliling persegi panjang: %.2f\n", keliling)
}
