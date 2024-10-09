package main

import (
	"fmt"
)

// Fungsi f(x) = x^2
func fungsiKuadrat(nilai int) int {
	return nilai * nilai
}

// Fungsi g(x) = x - 2
func fungsiKurang(nilai int) int {
	return nilai - 2
}

// Fungsi h(x) = x + 1
func fungsiTambah(nilai int) int {
	return nilai + 1
}


// Fungsi komposisi fogoh(a)
func komposisiFogoh(nilai int) int {
	return fungsiKuadrat(fungsiKurang(fungsiTambah(nilai))) // f(g(h(x)))
}

// Fungsi komposisi gohof(b)
func komposisiGohof(nilai int) int {
	return fungsiKurang(fungsiTambah(fungsiKuadrat(nilai))) // g(h(f(x)))
}

// Fungsi komposisi hofog(c)
func komposisiHofog(nilai int) int {
	return fungsiTambah(fungsiKuadrat(fungsiKurang(nilai))) // h(f(g(x)))
}

func main() {
	// Meminta input dari pengguna
	var a, b, c int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Menghitung hasil komposisi
	hasilFogoh := komposisiFogoh(a)
	hasilGohof := komposisiGohof(b)
	hasilHofog := komposisiHofog(c)

	// Menampilkan hasil
	fmt.Printf("\nHasil komposisi (fogoh)(%d): %d\n", a, hasilFogoh)
	fmt.Printf("Hasil komposisi (gohof)(%d): %d\n", b, hasilGohof)
	fmt.Printf("Hasil komposisi (hofog)(%d): %d\n", c, hasilHofog)
}
