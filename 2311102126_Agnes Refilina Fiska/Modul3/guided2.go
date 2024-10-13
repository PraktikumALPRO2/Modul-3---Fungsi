package main

import "fmt"

func main() {
	var panjang, lebar int

	fmt.Print("Masukkan panjang persegi: ")
	fmt.Scanln(&panjang)

	fmt.Print("Masukkan lebar persegi: ")
	fmt.Scanln(&lebar)

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Printf("Luas persegi: %d\n", luas)
	fmt.Printf("Keliling persegi: %d\n", keliling)
}
