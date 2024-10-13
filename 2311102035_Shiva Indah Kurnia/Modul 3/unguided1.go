package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
func main() {
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan bulat: ")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutation(a, c)
	c1 := combination(a, c)
	fmt.Printf("%d %d\n", p1, c1)

	p2 := permutation(b, d)
	c2 := combination(b, d)
	fmt.Printf("%d %d\n", p2, c2)
}
