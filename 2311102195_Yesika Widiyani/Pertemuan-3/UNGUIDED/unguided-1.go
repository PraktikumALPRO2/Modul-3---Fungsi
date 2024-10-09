package main

import (
	"fmt"
)

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
	// Input bilangan a, b, c
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// Menghitung hasil komposisi
	result1 := fogoh(a)
	result2 := gohof(b)
	result3 := hofog(c)

	// Output hasil komposisi
	fmt.Println(result1) // (f o g o h)(a)
	fmt.Println(result2) // (g o h o f)(b)
	fmt.Println(result3) // (h o f o g)(c)
}
