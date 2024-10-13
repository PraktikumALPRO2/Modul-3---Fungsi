/* Liya Khoirunnisa - 2311102124*/

package main

import (
	"fmt"
)

// Fungsi untuk f(x)
func f(x int) int {
	return x * x
}

// Fungsi untuk g(x)
func g(x int) int {
	return x - 2
}

// Fungsi untuk h(x)
func h(x int) int {
	return x + 1
}

// Fungsi komposisi f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi komposisi g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi komposisi h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int
	// Membaca input dari pengguna
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Menghitung hasil komposisi fungsi
	fmt.Printf("\nHasil: \n")
	fmt.Printf("fogoh(%d) = %d\n", a, fogoh(a)) // (fogoh)(a)
	fmt.Printf("gohof(%d) = %d\n", b, gohof(b)) // (gohof)(b)
	fmt.Printf("hofog(%d) = %d\n", c, hofog(c)) // (hofog)(c)
}