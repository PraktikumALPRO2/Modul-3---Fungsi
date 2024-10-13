package main

import (
	"fmt"
)

// Definisi fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Definisi fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Definisi fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Komposisi f(g(h(x))) -> (fogoh)(x)
func fogoh(x int) int {
	return f(g(h(x)))
}

// Komposisi g(h(f(x))) -> (gohof)(x)
func gohof(x int) int {
	return g(h(f(x)))
}

// Komposisi h(f(g(x))) -> (hofog)(x)
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	// Input tiga bilangan: a, b, c
	var a, b, c int
	fmt.Print("Masukkan 3 bilangan (pisahkan dengan spasi): ")
	fmt.Scan(&a, &b, &c)

	// Output hasil dari komposisi fungsi
	fmt.Println(fogoh(a)) // Hasil (fogoh)(a)
	fmt.Println(gohof(b)) // Hasil (gohof)(b)
	fmt.Println(hofog(c)) // Hasil (hofog)(c)
}
