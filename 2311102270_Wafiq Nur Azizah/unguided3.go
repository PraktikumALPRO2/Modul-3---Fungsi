package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Println("Lingkaran 1")
	fmt.Print("(cx)(cy)(r) : ")
	fmt.Scanln(&cx1, &cy1, &r1)

	fmt.Println("Lingkaran 2")
	fmt.Print("(cx)(cy)(r) : ")
	fmt.Scanln(&cx2, &cy2, &r2)

	fmt.Print("Input koordinat sembarang (x)(y) : ")
	fmt.Scanln(&x, &y)

	// Cek posisi terhadap Lingkaran 1 dan Lingkaran 2
	isInsideCircle1 := isPointInCircle(cx1, cy1, x, y, r1)
	isInsideCircle2 := isPointInCircle(cx2, cy2, x, y, r2)

	if isInsideCircle1 && isInsideCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if isInsideCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if isInsideCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func calculateDistance(a, b, c, d float64) float64 {
	// Menghitung jarak antara dua titik
	dx := (a - c) * (a - c)
	dy := (b - d) * (b - d)
	return math.Sqrt(dx + dy)
}

func isPointInCircle(cx, cy, x, y, r float64) bool {
	// Mengecek apakah titik (x, y) berada dalam lingkaran
	return calculateDistance(cx, cy, x, y) <= r
}