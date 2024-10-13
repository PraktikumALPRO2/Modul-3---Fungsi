package main
import "fmt"

func luas (sisi int) int {
	luaspersegi := sisi * sisi
	return luaspersegi
}

func keliling (sisi int) int {
	kelilingpersegi := 4*sisi
	return kelilingpersegi
}

func main(){
	var sisi int
	fmt.Print("Masukkan Sisi Persegi: ")
	fmt.Scan(&sisi)
	fmt.Println("Persegi dengan sisi" ,sisi,"memiliki luas", luas(sisi), "dan keliling", keliling (sisi))
}