package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
	if n < r {
		return 0 // Tidak valid
	}
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int) int {
	if n < r {
		return 0 // Tidak valid
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan nilai a dan b (dengan a >= c dan b >= d): ")
	fmt.Scan(&a, &b)
	fmt.Print("Masukkan nilai c dan d: ")
	fmt.Scan(&c, &d)

	if a >= c && b >= d {
		// Baris pertama: Hasil permutasi dan kombinasi a terhadap c
		fmt.Printf("Permutasi P(%d, %d) = %d\n", a, c, permutasi(a, c))
		fmt.Printf("Kombinasi C(%d, %d) = %d\n", a, c, kombinasi(a, c))

		// Baris kedua: Hasil permutasi dan kombinasi b terhadap d
		fmt.Printf("Permutasi P(%d, %d) = %d\n", b, d, permutasi(b, d))
		fmt.Printf("Kombinasi C(%d, %d) = %d\n", b, d, kombinasi(b, d))
	} else {
		fmt.Println("Pastikan bahwa a >= c dan b >= d.")
	}
}