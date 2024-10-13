package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Print("Input nilai a, b, c, d : ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println("Hasil Permutasi a: ", hitungPermutasi(a, c))
		fmt.Println("Hasil Kombinasi a: ", hitungKombinasi(a, c))
		fmt.Println("Hasil Permutasi b: ", hitungPermutasi(b, d))
		fmt.Println("Hasil Kombinasi b: ", hitungKombinasi(b, d))
	} else {
		fmt.Print("Kondisi 'a >= c && b >= d' tidak terpenuhi: ", a, b, c, d)
	}
}

func hitungFaktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func hitungPermutasi(n, r int) int {
	return hitungFaktorial(n) / hitungFaktorial(n-r)
}

func hitungKombinasi(n, r int) int {
	return hitungFaktorial(n) / (hitungFaktorial(r) * hitungFaktorial(n-r))
}