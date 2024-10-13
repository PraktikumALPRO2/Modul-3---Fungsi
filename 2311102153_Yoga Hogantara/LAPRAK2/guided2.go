package main

import "fmt"

func luas(s int) int {
	var hasil int
	hasil=s*s
	return hasil
}
func keliling(s int) int {
	var hasil int
	hasil=s*4
	return hasil
}

func main(){
	var s int
	fmt.Print("SISI:")
	fmt.Scan(&s)
	fmt.Println("Luas: ",luas(s))
	fmt.Println("keliling: ",keliling(s))
	
}
