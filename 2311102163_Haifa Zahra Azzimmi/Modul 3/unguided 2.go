package main

import (
	"fmt"
)

// Definisikan fungsi f(x), g(x), h(x)
func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

// Fungsi komposisi untuk fogoh(x), gohof(x), dan hofog(x)
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
	var a, b, c int

	// Input bilangan bulat a, b, c
	fmt.Print("Masukkan nilai a, b, c (dipisahkan dengan spasi): ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	// Output dengan keterangan
	fmt.Printf("Hasil fogoh(%d) = %d\n", a, fogoh(a))    // fogoh(a)
	fmt.Printf("Hasil gohof(%d) = %d\n", b, gohof(b))    // gohof(b)
	fmt.Printf("Hasil hofog(%d) = %d\n", c, hofog(c))    // hofog(c)
}
