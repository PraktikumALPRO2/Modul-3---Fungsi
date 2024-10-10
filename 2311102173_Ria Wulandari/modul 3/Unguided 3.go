package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func diDalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	// Input untuk lingkaran 1
	fmt.Println("Masukkan koordinat pusat (cx, cy) dan radius r lingkaran 1:")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input untuk lingkaran 2
	fmt.Println("Masukkan koordinat pusat (cx, cy) dan radius r lingkaran 2:")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input untuk titik sembarang
	fmt.Println("Masukkan koordinat titik sembarang (x, y):")
	fmt.Scan(&x, &y)

	// Mengecek posisi titik terhadap lingkaran 1 dan lingkaran 2
	dalamLingkaran1 := diDalamLingkaran(cx1, cy1, r1, x, y)
	dalamLingkaran2 := diDalamLingkaran(cx2, cy2, r2, x, y)

	// Menentukan keluaran berdasarkan posisi titik
	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
