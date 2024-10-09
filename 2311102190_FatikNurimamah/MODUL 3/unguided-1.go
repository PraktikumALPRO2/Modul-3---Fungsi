package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func HitungFaktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func HitungPermutasi(n, r int) int {
	if n < r {
		return 0
	}
	return HitungFaktorial(n) / HitungFaktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func HitungKombinasi(n, r int) int {
	if n < r {
		return 0
	}
	return HitungFaktorial(n) / (HitungFaktorial(r) * HitungFaktorial(n-r))
}


func main() {
	var a, b, c, d int

	// Meminta pengguna untuk memasukkan input
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	fmt.Print("Masukkan nilai d: ")
	fmt.Scan(&d)

	// Menghitung permutasi dan kombinasi
	hasilPermutasi1 := HitungPermutasi(a, c)
	hasilKombinasi1 := HitungKombinasi(a, c)
	hasilPermutasi2 := HitungPermutasi(b, d)
	hasilKombinasi2 := HitungKombinasi(b, d)

	// Menampilkan hasil
	fmt.Printf("\nPermutasi dan Kombinasi untuk %d dan %d: %d %d\n", a, c, hasilPermutasi1, hasilKombinasi1)
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", b, d, hasilPermutasi2, hasilKombinasi2)
}
