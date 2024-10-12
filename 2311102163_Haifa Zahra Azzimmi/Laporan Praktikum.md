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
![image](https://github.com/user-attachments/assets/ecd2009f-4cf7-4de1-93d0-1a3cc8a298ee)

#### Full Code Screenshot :
![image](https://github.com/user-attachments/assets/04a456f8-34f6-46c6-8715-f063cc15397c)

#### Deskripsi Program :
Program ini digunakan untuk menghitung permutasi dari dua angka yang diinputksn oleh pengguna. Setelah dua angka dimasukkan dan disimpan dalam variabel a dan b, program memeriksa mana yang lebih besar. Jika a lebih besar atau sama dengan b, maka program akan menghitung permutasi P (a, b), tetapi jika b lebih besar dari a, permutasi yang dihitung adalah P (b, a).

Program ini memungkinkan pengguna untuk menghitung berapa banyak cara elemen-elemen dapat disusun dalam jumlah tertentu dari total elemen yang tersedia.

#### Algoritma Program
1. Mulai
2. Input dua angka a dan b.
3. Jika b lebih besar dari a, tukar nilainya.
4. Panggil fungsi permutasi (a, b) yang menghitung P (a, b) dengan:
   - Menggunakan fungsi faktorial untuk menghitung faktorial a dan (a - b).
   - Bagikan hasil faktorial a dengan faktorial (a - b).
5. Tampilkan hasil permutasi.
6. Selesai
7. 
#### Cara Kerja Program
1. Fungsi faktorial menerima sebuah bilangan bulat dan menghitung faktorialnya dengan mengalikan semua angka dari 1 hingga bilangan tersebut.
2. Hasil faktorial ini kemudian digunakan oleh fungsi permutasi untuk menghitung nilai permutasi berdasarkan rumus yang sesuai.
3. Program akan menampilkan hasil akhir dengan membagi faktorial dari `n` dengan faktorial dari `n - r` untuk mendapatkan nilai permutasi.


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

#### Deskripsi Program	:
Program ini digunakan untuk menghitung luas dan keliling persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna. Setelah menerima input, luas dihitung dengan mengalikan panjang sisi dengan dirinya sendiri, sementara keliling dihitung dengan mengalikan panjang sisi dengan 4. Hasilnya ditampilkan menggunakan `fmt.Printf()`. Program ini sederhana dan efektif untuk memberikan informasi tentang persegi.

#### Algoritma Program :
1. Mulai
2. Deklarasikan variabel panjangSisi dengan tipe float64.
3. Tampilkan instruksi kepada pengguna untuk memasukkan panjang sisi persegi.
4. Ambil input dari pengguna dan simpan dalam panjangSisi.
5. Hitung luas persegi:
   - LuasPersegi = panjangSisi * panjangSisi
6. Hitung keliling persegi:
   - Keliling Persegi = 4 * panjangSisi
7. Tampilkan hasil perhitungan luas dan keliling di layar.
8. Selesai

#### Cara Kerja Program	:
1. Program meminta pengguna untuk memasukkan panjang sisi persegi dan menyimpannya dalam variabel panjang sisi.
2. Program menghitung luas dengan rumus Luas
3. Menghitung Keliling: Program menghitung keliling menggunakan rumus Keliling
4. Hasil luas dan keliling ditampilkan di layar menggunakan fmt.Printf().


## Unguided
### 1. Program untuk menghitung Faktorial, Permutasi dan Kombinasi
### Source Code :

### Output :

### Full code Screenshot :

### Algoritma Program :

### Cara Kerja Program :



### 2. Program untuk menghitung Faktorial, Permutasi dan Kombinasi
### Source Code :

### Output :

### Full code Screenshot :

### Algoritma Program :

### Cara Kerja Program :



### 3. Program untuk menghitung Faktorial, Permutasi dan Kombinasi
### Source Code :

### Output :


### Full code Screenshot :

### Algoritma Program :

### Cara Kerja Program :

