package main

import "fmt"

func f(x_145 int) int {
	return x_145 * x_145
}

func g(x_145 int) int {
	return x_145 - 2
}

func h(x_145 int) int {
	return x_145 + 1
}

func fogoh(x_145 int) int {
	return f(g(h(x_145)))
}

func gohof(x_145 int) int {
	return g(h(f(x_145)))
}

func hofog(x_145 int) int {
	return h(f(g(x_145)))
}

func main() {
	var a, b, c int
	fmt.Println("Masukkan tiga bilangan bulat (a, b, c):")
	fmt.Scanln(&a, &b, &c)

	fmt.Println("(fogoh)(a) =", fogoh(a))
	fmt.Println("(gohof)(b) =", gohof(b))
	fmt.Println("(hofog)(c) =", hofog(c))
}
