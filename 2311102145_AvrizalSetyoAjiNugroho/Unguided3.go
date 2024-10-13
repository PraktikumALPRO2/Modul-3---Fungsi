package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (cx1, cy1, r1):")
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 float64
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (cx2, cy2, r2):")
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y float64
	fmt.Println("Masukkan koordinat titik (x, y):")
	fmt.Scan(&x, &y)

	inCircle1 := didalam(cx1, cy1, r1, x, y)
	inCircle2 := didalam(cx2, cy2, r2, x, y)

	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
