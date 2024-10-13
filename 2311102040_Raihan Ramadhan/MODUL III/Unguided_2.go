package main

import (
	"fmt"
)

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

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
	fmt.Print("Masukkan nilai a, b, dan c: ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	result1 := fogoh(a)
	result2 := gohof(b)
	result3 := hofog(c)

	fmt.Printf("(fogoh)(%d) = %d\n", a, result1)
	fmt.Printf("(gohof)(%d) = %d\n", b, result2)
	fmt.Printf("(hofog)(%d) = %d\n", c, result3)
}
