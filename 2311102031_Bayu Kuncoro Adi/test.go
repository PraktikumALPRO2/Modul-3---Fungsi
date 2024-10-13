package main

import (
	"fmt"
)

func main() {
	var sisi float64

	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scanln(&sisi)

	luas := sisi * sisi

	keliling := 4 * sisi

	fmt.Printf("Luas persegi dengan sisi %.2f adalah %.2f\n", sisi, luas)
	fmt.Printf("Keliling persegi dengan sisi %.2f adalah %.2f\n", sisi, keliling)
}
