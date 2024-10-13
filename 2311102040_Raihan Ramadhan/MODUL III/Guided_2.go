package main

import "fmt"

var sisi int

func main() {
	fmt.Print("Masukan Panjang Sisi: ")
	fmt.Scan(&sisi)
	luas := LuasPersegi(sisi)
	keliling := KelilingPersegi(sisi)
	fmt.Printf("Luas Persegi adalah: %d\n", luas)
	fmt.Printf("Keliling Persegi adalah: %d\n", keliling)
}

func LuasPersegi(sisi int) int {
	return sisi * sisi
}

func KelilingPersegi(sisi int) int {
	return 4 * sisi
}
