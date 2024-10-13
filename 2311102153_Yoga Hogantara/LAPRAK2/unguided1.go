package main

import "fmt"

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i =1; i<=n; i++{
		hasil= hasil*i
	}
	return hasil
}
func permutasi(n,r int) int {
	return faktorial(n)/faktorial(n-r)
}
func kombinasi(n, r int)int{
	return faktorial(n) / (faktorial(r) * faktorial(n-r))

}
func main(){
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan: ")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	c1 := kombinasi(a, c)
	fmt.Printf("%d %d\n", p1, c1) 
	p2 := permutasi(b, d)
	c2 := kombinasi(b, d)
	fmt.Printf("%d %d\n", p2, c2)


}