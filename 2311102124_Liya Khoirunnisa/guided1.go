package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c && b >= d {
		// Hitung permutasi dan kombinasi untuk a dan c
		fmt.Println(permutasi(a, c), kombinasi(a, c))
		// Hitung permutasi dan kombinasi untuk b dan d
		fmt.Println(permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println("Input tidak valid, pastikan a >= c dan b >= d")
	}
}

func faktorial(n int) int {
	hasil := 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
