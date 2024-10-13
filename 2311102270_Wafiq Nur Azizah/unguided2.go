package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Input a, b, c : ")
	fmt.Scanln(&a, &b, &c)

	hasil1 := komposisiFogoh(a)
	hasil2 := komposisiGohof(b)
	hasil3 := komposisiHofog(c)
	fmt.Printf("(fogoh)(%d) = %d\n", a, hasil1)
	fmt.Printf("(gohof)(%d) = %d\n", b, hasil2)
	fmt.Printf("(hofog)(%d) = %d\n", c, hasil3)
}

// Definisi fungsi matematika
// f(x) := x * x
// g(x) := x - 2
// h(x) := x + 1

func komposisiFogoh(x int) int {
	// f(g(h(x)))
	hx := h(x)
	ghx := g(hx)
	result := f(ghx)

	return result
}

func komposisiGohof(x int) int {
	// g(h(f(x)))
	fx := f(x)
	hfx := h(fx)
	result := g(hfx)

	return result
}

func komposisiHofog(x int) int {
	// h(f(g(x)))
	gx := g(x)
	fgx := f(gx)
	result := h(fgx)

	return result
}

// Definisi fungsi f, g, dan h
func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}