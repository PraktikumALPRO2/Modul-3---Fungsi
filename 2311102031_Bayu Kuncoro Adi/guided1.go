package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b) // Membaca input dari pengguna untuk dua nilai integer a dan b
	if a >= b {
		// Jika a lebih besar atau sama dengan b, panggil fungsi permutasi dengan parameter (a, b)
		fmt.Println(permutasi(a, b))
	} else {
		// Jika b lebih besar dari a, panggil fungsi permutasi dengan parameter (b, a)
		fmt.Println(permutasi(b, a))
	}
}

func faktorial(n int) int {
	var hasil int = 1
	var i int
	// Loop untuk menghitung faktorial dari n
	for i = 1; i <= n; i++ {
		hasil = hasil * i // Mengalikan hasil dengan nilai i pada setiap iterasi
	}
	return hasil // Mengembalikan hasil faktorial
}

func permutasi(n, r int) int {
	// Menghitung permutasi nPr dengan membagi faktorial n dengan faktorial (n-r)
	return faktorial(n) / faktorial(n-r)
}
