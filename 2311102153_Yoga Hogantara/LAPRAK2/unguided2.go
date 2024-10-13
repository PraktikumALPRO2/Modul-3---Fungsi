package main

import (
	"fmt"
)

func f(x int) int {
	var f int
	f = x*x
	return f
}

func g(x int) int {
	var g int
	g = x-2
	return g
}

func h(x int) int {
	var h int
	h = x+1
	return h
}

func fgh(x int) int {
	var fogoh int
	fogoh = f(g(h(x)))
	return fogoh
}

func ghf(x int) int {
	var gohof int
	gohof = g(h(f(x)))
	return gohof
}

func hfg(x int) int {
	var hofog int
	hofog = h(f(g(x)))
	return hofog
}

func main() {
	var a, b, c int
	fmt.Print("INPUT a,b,c : \n")
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println("fogoh(a) =", fgh(a))
	fmt.Println("gohof(b) =", ghf(b))
	fmt.Println("hofog(c) =", hfg(c))
}
