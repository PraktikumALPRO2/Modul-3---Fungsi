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
  Arif Amrulloh
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
### 1. Struktur Program G0
Bahasa pemrograman Go (sering disebut Golang) memiliki struktur yang sederhana dan efisien. Berikut adalah komponen utama dalam struktur program Go:

**1. `Package Declaration`**
Semua program Go dimulai dengan deklarasi `package`. Program yang dapat dieksekusi (standalone) menggunakan `package main`.

```go
package main
```

**2. `Import Statement`**
Digunakan untuk memasukkan package atau pustaka standar/buatan yang akan digunakan dalam program.

```go
import "fmt"
```

**3. `Function Declaration`**
Fungsi dalam Go dideklarasikan menggunakan kata kunci `func`. Fungsi `main()` adalah fungsi yang dieksekusi pertama kali dalam program Go.

```go
func main() {
    fmt.Println("Hello, World!")
}
```

### Tipe Data Dalam GO
Go memiliki beberapa tipe data dasar yang digunakan untuk menyimpan berbagai jenis nilai, seperti bilangan, karakter, string, atau nilai boolean.

#### A. Tipe Data Integer
Tipe data integer digunakan untuk menyimpan bilangan bulat (tanpa desimal). Go mendukung dua jenis bilangan bulat: signed (bertanda) dan unsigned (tidak bertanda).

**1. `Signed Integer :`**
Mendukung bilangan positif dan negatif.

`int8`: Berisi nilai dari -128 hingga 127.

`int16`: Berisi nilai dari -32,768 hingga 32,767.

`int32`: Berisi nilai dari -2,147,483,648 hingga 2,147,483,647.

`int64`: Berisi nilai dari -9,223,372,036,854,775,808 hingga 9,223,372,036,854,775,807.

`int`: Ukurannya tergantung arsitektur komputer (32-bit atau 64-bit).

**2. `Unsigned Integer :`**
Hanya mendukung bilangan positif.

`uint8`: Berisi nilai dari 0 hingga 255.

`uint16`: Berisi nilai dari 0 hingga 65,535.

`uint32`: Berisi nilai dari 0 hingga 4,294,967,295.

`uint64`: Berisi nilai dari 0 hingga 18,446,744,073,709,551,615.

`uint`: Ukurannya tergantung arsitektur komputer (32-bit atau 64-bit).

#### B. Tipe Data Floating-Point (Real)
Tipe floating-point digunakan untuk menyimpan bilangan desimal (pecahan).

`float32`: Mendukung angka pecahan dengan presisi hingga 6-7 digit.

`float64`: Mendukung angka pecahan dengan presisi hingga 15 digit.

contoh :
```go
var x float32 = 3.14
var y float64 = 123.456789
```

#### C. Tipe Data Boolean
Tipe boolean digunakan untuk menyimpan nilai logika, yaitu `true` (benar) atau `false` (salah). Biasanya digunakan dalam pengambilan keputusan (kondisi).

contoh:
```go
var isLearningGo bool = true
```

#### D. Tipe Data String
Tipe string digunakan untuk menyimpan serangkaian karakter (teks). Dalam Go, string dideklarasikan menggunakan tanda kutip ganda (`"`). String bersifat immutable, yang berarti nilainya tidak bisa diubah setelah dideklarasikan.

contoh:
```go
var greeting string = "Selamat Datang di Golang"
```

#### E. Tipe Data Karakter
Go tidak memiliki tipe data khusus untuk karakter. Sebagai gantinya, karakter ditangani sebagai byte atau rune:

`byte`: Alias untuk `uint8`, digunakan untuk merepresentasikan karakter ASCII.
`rune`: Alias untuk `int32`, digunakan untuk karakter Unicode.

contoh:
```go
var ch rune = 'A'
fmt.Println(ch)  // Output: 65 (kode Unicode/ASCII untuk 'A')
```

### Instruksi Dasar Dalam GO
#### A. Deklarasi Dan Inisialisasi Variabel
**Menggunakan `var` dan inisialisasi manual:**
```go
var x int
x = 10
```
**Deklarasi dan inisialisasi secara langsung:**
```go
var y int = 20
```
**Deklarasi singkat menggunakan `:=`:**
```go
z := 30
```
#### B. Input dan Output
Fungsi input/output utama dalam Go terdapat pada package `fmt`.

**Output (Print ke layar):**
```go
fmt.Println("Halo, Go!")
fmt.Printf("Nilai a adalah: %d\n", a)  // Format output dengan placeholder
```
**Input (Membaca dari pengguna):**
```go
var umur int
fmt.Print("Masukkan umur Anda: ")
fmt.Scanln(&umur)  // Membaca input dari pengguna
```
#### C. Struktur Kontrol
Go mendukung berbagai instruksi untuk mengontrol alur program seperti:

**1. `If-else` (Kondisional):**
```go
if umur >= 18 {
    fmt.Println("Anda sudah dewasa.")
} else {
    fmt.Println("Anda masih anak-anak.")
}
```
**2. `Switch` (Pemilihan multi-cabang):**
```go
switch umur {
case 18:
    fmt.Println("Anda baru saja dewasa!")
case 30:
    fmt.Println("Selamat datang di usia 30!")
default:
    fmt.Println("Umur Anda adalah:", umur)
}
```
**3. `For Loop` (Perulangan):**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## Guided
### 1. 
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

#### Deskripsi Program :

#### Algoritma Program

#### Cara Kerja Program

### 2. 
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

#### Deskripsi Program

#### Algoritma Program :


#### Cara Kerja Program



