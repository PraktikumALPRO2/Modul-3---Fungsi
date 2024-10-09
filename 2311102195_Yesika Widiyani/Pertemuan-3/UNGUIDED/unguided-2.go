package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func hitungJarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

// Fungsi untuk mengecek apakah titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan radius r
func dalamLingkaran(x, y, cx, cy, r int) bool {
	return hitungJarak(x, y, cx, cy) <= float64(r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	// Input untuk lingkaran 1
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input untuk lingkaran 2
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input untuk titik sembarang
	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scan(&x, &y)

	// Cek posisi titik
	dalam1 := dalamLingkaran(x, y, cx1, cy1, r1)
	dalam2 := dalamLingkaran(x, y, cx2, cy2, r2)

	// Output hasil berdasarkan posisi titik
	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
