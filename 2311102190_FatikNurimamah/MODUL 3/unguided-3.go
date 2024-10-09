package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func hitungJarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

// Fungsi untuk mengecek apakah suatu titik (x, y) berada di dalam lingkaran
func cekdiDalamLingkaran(pusatX, pusatY, radius, titikX, titikY int) bool {
	return hitungJarak(pusatX, pusatY, titikX, titikY) <= float64(radius)
}

func main() {
	// Membaca input
	var pusatX1, pusatY1, radius1 int // Koordinat dan radius lingkaran 1
	var pusatX2, pusatY2, radius2 int // Koordinat dan radius lingkaran 2
	var titikX, titikY int            // Koordinat titik sembarang

	fmt.Print("Masukkan koordinat dan radius lingkaran 1: ")
	fmt.Scanln(&pusatX1, &pusatY1, &radius1)

	fmt.Print("Masukkan koordinat dan radius lingkaran 2: ")
	fmt.Scanln(&pusatX2, &pusatY2, &radius2)

	fmt.Print("Masukkan koordinat titik: ")
	fmt.Scanln(&titikX, &titikY)

	// Mengecek posisi titik (titikX, titikY) terhadap lingkaran 1 dan lingkaran 2
	diLingkaran1 := cekdiDalamLingkaran(pusatX1, pusatY1, radius1, titikX, titikY)
	diLingkaran2 := cekdiDalamLingkaran(pusatX2, pusatY2, radius2, titikX, titikY)

	// Menentukan output berdasarkan posisi titik
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("\nTitik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("\nTitik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("\nTitik di dalam lingkaran 2")
	} else {
		fmt.Println("\nTitik di luar lingkaran 1 dan 2")
	}
}
