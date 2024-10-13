package main

import (
	"fmt"
)


func Luas(s int) int{
	var L int
	L = s * s
	return L
}
func Keliling(s int) int{
	var K int 
	K = s+s+s+s
	return K
}
func main(){
	var s int
	fmt.Println("Inputkan sisi: ")
	fmt.Scan(&s)
	fmt.Println("Luas: ",Luas(s))
	fmt.Println("keliling: ",Keliling(s))
}