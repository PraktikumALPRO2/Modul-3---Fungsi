/*Liya Khoirunnisa - 2311102124*/

package main

import (
	"fmt"
	"math" // Untuk operasi matematika
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func cekDalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	// Deklarasi variabek
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	// Input lingkaran 1
	fmt.Println("Lingkaran 1")
	fmt.Print("Masukkan titik pusat (cx1, cy1) dan radius r1: ")
	fmt.Scan(&cx1, &cy1, &r1)
	
	// Input lingkaran 2
	fmt.Println("\nLingkaran 2")
	fmt.Print("Masukkan titik pusat (cx2, cy2) dan radius r2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input titik sembarang
	fmt.Print("\nMasukkan titik sembarang (x, y): ")
	fmt.Scan(&x, &y)

	// Mengecek apakah titik berada di dalam lingkaran 1 dan/atau lingkaran 2
	titikDalam1 := cekDalamLingkaran(cx1, cy1, r1, x, y)
	titikDalam2 := cekDalamLingkaran(cx2, cy2, r2, x, y)
	
	// Menentukan output berdasarkan posisi titik
	fmt.Println("\nPosisi titik")
	if titikDalam1 && titikDalam2 {
		// Jika titik berada di dalam kedua lingkaran
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if titikDalam1 {
		// Jika titik hanya berada di dalam ligkaran 1
		fmt.Println("Titik di dalam lingkaran 1")
	} else if titikDalam2 {
		// Jika titik hanya berada di lingkaran 2
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		// Jika titik berada di luar kedua lingkaran
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
