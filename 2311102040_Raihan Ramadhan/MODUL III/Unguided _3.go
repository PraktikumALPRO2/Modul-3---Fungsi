package main

import (
	"fmt"
	"math"
)

func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Hypot(x2-x1, y2-y1)
}

func cekDalamLingkaran(cx, cy, radius, px, py float64) bool {
	return hitungJarak(cx, cy, px, py) <= radius
}

func main() {
	var pusatX1, pusatY1, radius1 float64
	var pusatX2, pusatY2, radius2 float64
	var koordinatX, koordinatY float64

	fmt.Println("Masukkan pusatX1, pusatY1, radius1 untuk Lingkaran 1: ")
	fmt.Scan(&pusatX1, &pusatY1, &radius1)
	fmt.Println("Masukkan pusatX2, pusatY2, radius2 untuk Lingkaran 2: ")
	fmt.Scan(&pusatX2, &pusatY2, &radius2)
	fmt.Println("Masukkan koordinat titik x dan y: ")
	fmt.Scan(&koordinatX, &koordinatY)

	adaDiLingkaran1 := cekDalamLingkaran(pusatX1, pusatY1, radius1, koordinatX, koordinatY)
	adaDiLingkaran2 := cekDalamLingkaran(pusatX2, pusatY2, radius2, koordinatX, koordinatY)

	if adaDiLingkaran1 && adaDiLingkaran2 {
		fmt.Println("Titik berada di dalam kedua lingkaran")
	} else if adaDiLingkaran1 {
		fmt.Println("Titik berada di dalam Lingkaran 1")
	} else if adaDiLingkaran2 {
		fmt.Println("Titik berada di dalam Lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar kedua lingkaran")
	}
}
