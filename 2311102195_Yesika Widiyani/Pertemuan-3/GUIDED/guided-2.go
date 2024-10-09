package main

import "fmt"

func main() {
	var sisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&sisi)

	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Printf("Luas persegi adalah: %.2f\n", luas)
	fmt.Printf("Keliling persegi adalah: %.2f\n", keliling)
}
