package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalamo(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64 
	var cx2, cy2, r2 float64 
	var x, y float64         

	fmt.Println("Lingkaran(O) 1 ")
	fmt.Println("INPUT cx1, cy1, r1 : ")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Println("Lingkaran(O) 2 ")
	fmt.Println("INPUT cx2, cy2, r2 : ")
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Println("INPUT(O) x, y: ")
	fmt.Scan(&x, &y)

	dio1 := didalamo(cx1, cy1, r1, x, y)
	dio2 := didalamo(cx2, cy2, r2, x, y)

	if dio1 && dio2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dio1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dio2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

