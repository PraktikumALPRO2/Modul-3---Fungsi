package main

import (
	"fmt"
)

//untuk menghitung faktorial
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
//untuk menghitung permutasi
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)

}
//untuk menghitung kombinasi
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	//input empat bilangan : a, b, c, d
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan : ")
	fmt.Scan(&a, &b, &c, &d)
	//baris pertama : permutasi dan kombinasi a terhadap c
	p1 := permutation(a, c)
	c1 := combination(a, c)
	fmt.Printf("%d %d\n", p1, c1)
	//bariskedua: permutasi dan kombinasi b terhadap d
	p2 := permutation(b, d)
	c2 := combination(b, d)
	fmt.Printf("%d %d\n", p2, c2)
	}