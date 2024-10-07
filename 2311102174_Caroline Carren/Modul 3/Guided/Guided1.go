// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi utama program
func main() {
	var a, b int
	// Menerima input dari pengguna
	fmt.Scan(&a, &b)
	// Memeriksa apakah a >= b
	if a >= b {
		// Jika ya, panggil fungsi permutasi(a, b)
		fmt.Println(permutasi(a, b))
	} else {
		// Jika tidak, panggil fungsi permutasi(b, a)
		fmt.Println(permutasi(b, a))
	}
}

// Fungsi untuk menghitung faktorial dari n
func faktorial(n int) int {
	var hasil int = 1
	var i int
	// Loop untuk menghitung faktorial
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi P(n, r)
func permutasi(n, r int) int {
	// Rumus permutasi: P(n, r) = n! / (n-r)!
	return faktorial(n) / faktorial(n-r)
}
