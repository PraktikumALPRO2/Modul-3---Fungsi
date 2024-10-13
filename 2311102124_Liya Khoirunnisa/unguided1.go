/* Liya Khoirunnisa - 2311102124 */

package main

import "fmt"

func main() {
	// Deklarasi variabel
	var a, b, c, d int

	// Input nilai a
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	// Input nilai b
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	// Input nilai c
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Input nilai d
	fmt.Print("Masukkan nilai d: ")
	fmt.Scan(&d)

	// Cek inputan
	if a >= c && b >= d {
		// Hitung permutasi dan kombinasi untuk a dan c
		fmt.Printf("\nP(%d,%d) = %d!/(%d)! = %d/%d = %d\n", a, c, a, a-c, faktorial(a), faktorial(a-c), permutasi(a, c))
		fmt.Printf("C(%d,%d) = %d!/(%d!x%d!) = %d/(%d x %d) = %d\n", a, c, a, c, a-c, faktorial(a), faktorial(c), faktorial(a-c), kombinasi(a, c))

		// Hitung permutasi dan kombinasi untuk b dan d
		fmt.Printf("P(%d,%d) = %d!/(%d)! = %d/%d = %d\n", b, d, b, b-d, faktorial(b), faktorial(b-d), permutasi(b, d))
		fmt.Printf("C(%d,%d) = %d!/(%d!x%d!) = %d/(%d x %d) = %d\n", b, d, b, d, b-d, faktorial(b), faktorial(d), faktorial(b-d), kombinasi(b, d))
	} else {
		fmt.Println("Input tidak valid, inputam harus a >= c dan b >= d")
	}
}

// Fungsi faktorial
func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

// Fungsi permutasi
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi kombinasi
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}