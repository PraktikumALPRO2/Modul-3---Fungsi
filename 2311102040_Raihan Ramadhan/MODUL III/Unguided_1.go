package main

import (
	"fmt"
)

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan nilai a, b, c, d : ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	permutasi_a_c := permutasi(a, c)
	kombinasi_a_c := kombinasi(a, c)

	permutasi_b_d := permutasi(b, d)
	kombinasi_b_d := kombinasi(b, d)

	fmt.Printf("Hasil permutasi dan kombinasi a terhadap c: %d %d\n", permutasi_a_c, kombinasi_a_c)
	fmt.Printf("Hasil permutasi dan kombinasi b terhadap d: %d %d\n", permutasi_b_d, kombinasi_b_d)
}
