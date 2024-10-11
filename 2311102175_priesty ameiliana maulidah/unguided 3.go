package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) < r
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scanf("%d %d %d", &cx1, &cy1, &r1)
	fmt.Scanf("%d %d %d", &cx2, &cy2, &r2)
	fmt.Scanf("%d %d", &x, &y)

	inLingkaran1 := didalam(float64(cx1), float64(cy1), float64(r1), float64(x), float64(y))
	inLingkaran2 := didalam(float64(cx2), float64(cy2), float64(r2), float64(x), float64(y))

	if inLingkaran1 && inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
