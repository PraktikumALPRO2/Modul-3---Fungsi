package main

import "fmt"

func factorial(n_145 int) int {
	if n_145 == 0 {
		return 1
	}
	return n_145 * factorial(n_145-1)
}

func permutation(n_145, r int) int {
	if n_145 < r {
		return 0
	}
	return factorial(n_145) / factorial(n_145-r)
}

func combination(n_145, r int) int {
	if n_145 < r {
		return 0
	}
	return factorial(n_145) / (factorial(r) * factorial(n_145-r))
}

func main() {
	var a, b, c, d int
	fmt.Println("Masukkan empat bilangan asli a, b, c, dan d, dengan syarat a >= c dan b >= d:")
	fmt.Scanln(&a, &b, &c, &d)

	p1 := permutation(a, c)
	c1 := combination(a, c)
	p2 := permutation(b, d)
	c2 := combination(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}