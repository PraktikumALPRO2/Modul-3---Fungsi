package main

import "fmt"

// Fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Fungsi h(x) = x + 1 
func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int 
	fmt.Print("Masukkan nilai a, b, c: ")
	fmt.Scan(&a, &b, &c)

	// Menghitung nilai a, b, c
	fogoh := f(g(h(a))) // (fogoh)(a)
	gohof := g(h(f(b))) // (gohof)(b)
	hofog := h(f(g(c))) // (hofog)(c)

	// Hasil perhitungan 
	fmt.Println("Hasil (fogoh)(a) : ", fogoh)
	fmt.Println("Hasil (gohof)(b) : ", gohof)
	fmt.Println("Hasil (hofog)(c) : ", hofog)
}