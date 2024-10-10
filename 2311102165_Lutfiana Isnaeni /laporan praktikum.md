<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL II</strong></h2>
<h2 align="center"><strong> FUNGSI </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Lutfiana Isnaeni L /2311102165<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh S.Kom., M.Kom. 
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

### 1 []

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
#### Output
![ss output guided1 modul3](https://github.com/user-attachments/assets/fa8fad50-9757-4bc4-a80e-0ff5d5407a66)

### Deskripsi Program
Kode di atas digumakan untuk menghitung nilai permutasi dari dua angka yang diinputkan oleh pengguna. Permutasi menghitung berapa banyak cara mengatur r objek dari n objek yang berbeda, dengan mempertimbangkan urutannya.

### Cara Kerja Program:
1. Fungsi main():
   
  - Program dimulai dengan membaca dua bilangan bulat, a dan b, dari input pengguna.

  - Jika a lebih besar atau sama dengan b, program akan menghitung permutasi P(a, b).

  - Jika sebaliknya, maka program menghitung permutasi P(b, a).

  - Hasil permutasi kemudian ditampilkan.

2. Fungsi faktorial(n int) int:
  - Fungsi ini digunakan untuk menghitung faktorial dari bilangan n.

  - Faktorial adalah hasil perkalian dari bilangan bulat positif dari 1 hingga n. Algoritma yang digunakan adalah iterasi, di mana program melakukan perkalian bertahap dari 1 sampai n.

  - Misalnya, faktorial dari 5 (5!) dihitung sebagai 1 * 2 * 3 * 4 * 5, yang hasilnya adalah 120.

3. Fungsi permutasi(n, r int) int:
  - Fungsi ini digunakan untuk menghitung permutasi P(n, r).

  - Rumus permutasi adalah: P(n, r) = n! / (n-r)!, yang berarti menghitung faktorial dari n kemudian membaginya dengan faktorial dari n - r.

  - Fungsi permutasi() memanggil fungsi faktorial() untuk melakukan perhitungan faktorial dari n dan n-r.

### Algoritma:
  - Input: Dua bilangan bulat a dan b dimasukkan oleh pengguna.

  - Proses: Program membandingkan nilai a dan b untuk menentukan urutan input bagi fungsi permutasi. Program menghitung permutasi menggunakan fungsi faktorial dengan iterasi.

- Output: Hasil dari perhitungan permutasi ditampilkan sebagai output.

### 2 []

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
#### Output
![ss output guided2](https://github.com/user-attachments/assets/06a31bb1-4382-42fa-b98a-c80fc76800aa)

### Deskripsi Program
Program di atas adalah sebuah program yang dibuat untuk menghitung luas dan keliling persegi berdasarkan input panjang sisi dari pengguna. Program ini mengikuti algoritma yang sederhana, yaitu menerima input, melakukan perhitungan luas dan keliling, lalu menampilkan hasilnya.

### Algoritma






   
