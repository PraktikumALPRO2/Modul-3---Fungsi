<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL III</strong></h2>
<h2 align="center"><strong> FUNGSI </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
 Haifa Zahra Azzimmi / 2311102163<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
Arif Amrulloh,S.Kom.,M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------


## Daftar Isi
1. [Dasar Teori](#dasar-teori)
2. [Guided](#guided)
3. [Unguided](#unguided)

## Dasar Teori


## Guided
### 1. Program fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi.
#### Source Code :
```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}
}
func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
```
#### Out Put :
![image](https://github.com/user-attachments/assets/a8147157-3fc8-4c27-8e11-5dc7e7167479)

#### Full Code Screenshot :
![image](https://github.com/user-attachments/assets/04a456f8-34f6-46c6-8715-f063cc15397c)

#### Deskripsi Program :
Program ini digunakan untuk menghitung permutasi dari dua angka yang diinputksn oleh pengguna. Setelah dua angka dimasukkan dan disimpan dalam variabel a dan b, program memeriksa mana yang lebih besar. Jika a lebih besar atau sama dengan b, maka program akan menghitung permutasi P(a, b), tetapi jika b lebih besar dari a, permutasi yang dihitung adalah P(b, a).

Program ini memungkinkan pengguna untuk menghitung berapa banyak cara elemen-elemen dapat disusun dalam jumlah tertentu dari total elemen yang tersedia.

#### Algoritma Program
1. Mulai
2. Input dua angka a dan b.
3. Jika b lebih besar dari a, tukar nilainya.
4. Panggil fungsi permutasi(a, b) yang menghitung P(a, b) dengan:
   - Menggunakan fungsi faktorial untuk menghitung faktorial a dan (a - b).
   - Bagikan hasil faktorial a dengan faktorial (a - b).
5. Tampilkan hasil permutasi.
6. Selesai
#### Cara Kerja Program

### 2. Program untuk menghitung luas persegi dan keliling persegi.
#### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	// Meminta input panjang sisi persegi
	var panjangSisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&panjangSisi)

	// Menghitung luas dan keliling persegi
	LuasPersegi := panjangSisi * panjangSisi
	KelilingPersegi := 4 * panjangSisi

	// Menampilkan hasil
	fmt.Printf("Luas persegi: %g\n", LuasPersegi)
	fmt.Printf("Keliling persegi: %g\n", KelilingPersegi)
}
```
#### Output :
![image](https://github.com/user-attachments/assets/6869e337-0427-45c2-8071-5412cf428b54)

#### Full Code Screenshot
![image](https://github.com/user-attachments/assets/f704df68-de88-40c7-99a6-c428373cebfb)

#### Deskripsi Program

#### Algoritma Program :


#### Cara Kerja Program



