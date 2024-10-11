package main

import (
	"fmt"
)

func main() {
	var sisi float64
	fmt.Print("masukkan panjang sisi persegi: ")
	fmt.Scan(&sisi)

	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Printf("luas persegi: %.2f\n", luas)
	fmt.Printf("keliling persegi: %.2f\n", keliling)
}