package main

import (
	"fmt"
)

// Definisikan fungsi f,g, dan h
func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

// Membuat fungsi komposisi
func fogoh(x int) int {
	return f(g(h(x)))
}

func gohof(x int) int {
	return g(h(f(x)))
}

func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	// Definisikan inputan
	var a, b, c int

	// Membaca Input
	fmt.Println("Masukkan 3 angka (a, b, c):")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	// Output
	fmt.Printf("f(g(h(%d))) = %d\n", a, fogoh(a))
	fmt.Printf("g(h(f(%d))) = %d\n", b, gohof(b))
	fmt.Printf("h(f(g(%d))) = %d\n", c, hofog(c))
}
