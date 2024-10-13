package main

import "fmt"

func luas_persegi(s int) int {
	luas := s * s
	return luas
}

func keliling_persegi(s int) int {

	keliling := 4 * s
	return keliling

}

func main() {
	var s int
	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scan(&s)
	fmt.Print("Jadi Luas persegi adalah: ", luas_persegi(s), "cm2 dan Keliling Persegi adalah: ", keliling_persegi(s), "cm")
}
