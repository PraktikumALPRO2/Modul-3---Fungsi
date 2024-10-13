package main

import "fmt"

// Fungsi untuk menentukan posisi titik terhadap sebuah lingkaran
func posisiTitikL(cx, cy, r, x, y int) bool {
	jarakKuadrat := (x-cx)*(x-cx) + (y-cy)*(y-cy) // Menghitung kuadrat jarak antara titik dan pusat lingkaran
	return jarakKuadrat <= r*r // Mengecek apakah jarak kuadrat titik dengan pusat lingkaran lebih lecil atau lebih besar
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	// Meminta pengguna untuk menginputkan titik pusat dan radius lingkaran 1
	fmt.Print("Masukkan cx1, cy1, r1 (ingkaran 1): ")
	fmt.Scan(&cx1, &cy2, &r1)

	// Meminta pengguna untuk menginputkan titik pusat dan radius lingkaran
	fmt.Print("Masukkan cx2, cy2, r2 (lingkaran 2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Meminta pengguna untuk menginputkan koordinat titik sembarang
	fmt.Print("Masukkan x, y (titik sembarang): ")
	fmt.Scan(&x, &y)

	// Mengecek posisi titik dalam lingkaran 1 & 2
	dalamL1 := posisiTitikL(cx1, cy1, r1, x, y)
	dalamL2 := posisiTitikL(cx2, cy2, r2, x, y)

	// Menampilkan output berdasarkan hasil posisi titik
	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}