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
  Fatik Nurimamah / 2311102190<br>
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
4. [Kesimpulan](#kesimpulan)
5. [Daftar Pustaka](#daftar-pustaka)


## Dasar Teori


## Guided 

### 1.Buatlah sebuah program beserta fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi

### Source Code :
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

### Output:


### Full code Screenshot:
![Screenshot 2024-10-09 085622](https://github.com/user-attachments/assets/335d5a58-0c57-4007-babe-9345028dd8d6)

### Deskripsi Program : 
Program meminta dua bilangan sebagai masukan, kemudian menghitung permutasi dengan memilih nilai terbesar sebagai jumlah total objek (n) dan nilai terkecil sebagai jumlah objek yang diambil (r). Program menggunakan fungsi faktorial untuk membantu perhitungan nilai permutasi tersebut.

### Algoritma Program
1. Program menerima dua angka bulat, a dan b, sebagai input.
2. Program kemudian membandingkan kedua angka tersebut untuk menentukan nilai terbesar sebagai n dan nilai terkecil sebagai r.
3. Fungsi permutasi(n, r) digunakan untuk menghitung nilai permutasi menggunakan rumus n!/(n - r)!, dengan ! sebagai operasi faktorial.
4. Fungsi faktorial dipanggil untuk menghitung faktorial dari n dan n - r.
5. Hasil perhitungan permutasi ditampilkan sebagai output.

### Cara Kerja Program:
1. Fungsi faktorial menerima bilangan bulat dan mengalikan semua angka dari 1 hingga angka tersebut untuk menghitung faktorialnya.
2. Fungsi permutasi kemudian menggunakan nilai faktorial yang diperoleh untuk menghitung permutasi sesuai rumus.
3. Program akhirnya menampilkan hasil perhitungan permutasi setelah membagi hasil faktorial dari n dengan faktorial dari n - r.
   
### 2. Program untuk menghitung luas persegi dan keliling persegi

### Source Code :
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

### Output:
![Screenshot 2024-10-09 090104](https://github.com/user-attachments/assets/fc3a1936-41c7-476b-8375-5f48aa345490)

### Full code Screenshot:
![Screenshot 2024-10-09 090155](https://github.com/user-attachments/assets/e53d16c7-1380-47ac-8145-62cdea270e15)

### Deskripsi Program : 
Program ini digunakan untuk menghitung luas dan keliling dari sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna. Program meminta pengguna memasukkan panjang sisi persegi, kemudian menghitung luas dan kelilingnya. Hasil dari kedua perhitungan tersebut kemudian ditampilkan sebagai output.

### Algoritma Program
1. Program meminta input panjang sisi persegi (panjangSisi) dari pengguna.
2. Menggunakan rumus untuk menghitung luas persegi, yaitu panjang sisi dikalikan dengan panjang sisi: LuasPersegi = panjangSisi * panjangSisi.
3. Menggunakan rumus untuk menghitung keliling persegi, yaitu panjang sisi dikalikan dengan 4: KelilingPersegi = 4 * panjangSisi.
4. Program menampilkan hasil perhitungan luas dan keliling persegi.

### Cara Kerja Program
1. Program menerima input panjang sisi persegi dan menyimpannya dalam variabel panjangSisi.
2. Luas persegi dihitung dengan mengalikan panjangSisi dengan dirinya sendiri.
3. Keliling persegi dihitung dengan mengalikan panjangSisi dengan 4.
4. Program kemudian menampilkan hasil luas dan keliling kepada pengguna.

## Unguided 

### 1. Program untuk menghitung Faktorial, Permutasi dan Kombinasi
![Screenshot 2024-10-09 090531](https://github.com/user-attachments/assets/939bb468-ffa3-42f2-a6a3-1e7d44633ed1)

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func HitungFaktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func HitungPermutasi(n, r int) int {
	if n < r {
		return 0
	}
	return HitungFaktorial(n) / HitungFaktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func HitungKombinasi(n, r int) int {
	if n < r {
		return 0
	}
	return HitungFaktorial(n) / (HitungFaktorial(r) * HitungFaktorial(n-r))
}


func main() {
	var a, b, c, d int

	// Meminta pengguna untuk memasukkan input
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	fmt.Print("Masukkan nilai d: ")
	fmt.Scan(&d)

	// Menghitung permutasi dan kombinasi
	hasilPermutasi1 := HitungPermutasi(a, c)
	hasilKombinasi1 := HitungKombinasi(a, c)
	hasilPermutasi2 := HitungPermutasi(b, d)
	hasilKombinasi2 := HitungKombinasi(b, d)

	// Menampilkan hasil
	fmt.Printf("\nPermutasi dan Kombinasi untuk %d dan %d: %d %d\n", a, c, hasilPermutasi1, hasilKombinasi1)
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", b, d, hasilPermutasi2, hasilKombinasi2)
}

```
### Output:
![Screenshot 2024-10-09 090843](https://github.com/user-attachments/assets/5f071abd-c8a3-4f53-a3d5-1c33eb557bbf)

### Full code Screenshot:
![Screenshot 2024-10-09 090935](https://github.com/user-attachments/assets/ac54511f-5101-446f-8952-fe9dab2635d3)

### Deskripsi Program : 
Program ini digunakan untuk menghitung nilai permutasi dan kombinasi dari dua pasang angka yang dimasukkan oleh pengguna. Fungsi HitungPermutasi berfungsi untuk menghitung jumlah pengurutan objek, sementara HitungKombinasi menghitung cara memilih objek tanpa memperhatikan urutan. Program meminta pengguna untuk memasukkan empat angka: a, b, c, dan d. Setelah itu, program menghitung permutasi dan kombinasi untuk pasangan (a, c) dan (b, d), lalu menampilkan hasilnya.

### Algoritma Program
1. Program meminta pengguna untuk menginput empat angka: a, b, c, dan d.
2. Fungsi HitungPermutasi digunakan untuk menghitung permutasi dengan rumus n!/(n - r)!.
3. Fungsi HitungKombinasi digunakan untuk menghitung kombinasi dengan rumus n!/r! (n - r)!.
4. Program menghitung permutasi dan kombinasi untuk pasangan (a, c) dan (b, d).
5. Hasil perhitungan permutasi dan kombinasi untuk kedua pasangan tersebut ditampilkan sebagai output.
   
### Cara Kerja Program
1. Program menerima input nilai a, b, c, dan d.
2. Untuk setiap pasangan (n, r), fungsi HitungPermutasi dipanggil untuk mendapatkan nilai permutasi dengan menghitung faktorial n dan n-r.
3. Fungsi HitungKombinasi juga dipanggil untuk setiap pasangan (n, r) menggunakan hasil faktorial untuk menghitung kombinasi.
4. Program kemudian menampilkan hasil permutasi dan kombinasi untuk pasangan (a, c) dan (b, d) di layar.

### 2.  
![Screenshot 2024-10-09 091205](https://github.com/user-attachments/assets/0613877e-afcf-448f-a46a-3fff064d9891)

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi f(x) = x^2
func fungsiKuadrat(nilai int) int {
	return nilai * nilai
}

// Fungsi g(x) = x - 2
func fungsiKurang(nilai int) int {
	return nilai - 2
}

// Fungsi h(x) = x + 1
func fungsiTambah(nilai int) int {
	return nilai + 1
}


// Fungsi komposisi fogoh(a)
func komposisiFogoh(nilai int) int {
	return fungsiKuadrat(fungsiKurang(fungsiTambah(nilai))) // f(g(h(x)))
}

// Fungsi komposisi gohof(b)
func komposisiGohof(nilai int) int {
	return fungsiKurang(fungsiTambah(fungsiKuadrat(nilai))) // g(h(f(x)))
}

// Fungsi komposisi hofog(c)
func komposisiHofog(nilai int) int {
	return fungsiTambah(fungsiKuadrat(fungsiKurang(nilai))) // h(f(g(x)))
}

func main() {
	// Meminta input dari pengguna
	var a, b, c int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Menghitung hasil komposisi
	hasilFogoh := komposisiFogoh(a)
	hasilGohof := komposisiGohof(b)
	hasilHofog := komposisiHofog(c)

	// Menampilkan hasil
	fmt.Printf("\nHasil komposisi (fogoh)(%d): %d\n", a, hasilFogoh)
	fmt.Printf("Hasil komposisi (gohof)(%d): %d\n", b, hasilGohof)
	fmt.Printf("Hasil komposisi (hofog)(%d): %d\n", c, hasilHofog)
}

```
### Output:
![Screenshot 2024-10-09 091414](https://github.com/user-attachments/assets/039e442d-eace-414a-b8db-3f24f986662a)

### Full code Screenshot:
![Screenshot 2024-10-09 092919](https://github.com/user-attachments/assets/93c480b6-37e9-452e-ad79-969708b37ebb)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program

### 3. Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut.
**Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".**


### Source Code :
```go
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func hitungJarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

// Fungsi untuk mengecek apakah suatu titik (x, y) berada di dalam lingkaran
func cekdiDalamLingkaran(pusatX, pusatY, radius, titikX, titikY int) bool {
	return hitungJarak(pusatX, pusatY, titikX, titikY) <= float64(radius)
}

func main() {
	// Membaca input
	var pusatX1, pusatY1, radius1 int // Koordinat dan radius lingkaran 1
	var pusatX2, pusatY2, radius2 int // Koordinat dan radius lingkaran 2
	var titikX, titikY int            // Koordinat titik sembarang

	fmt.Print("Masukkan koordinat dan radius lingkaran 1: ")
	fmt.Scanln(&pusatX1, &pusatY1, &radius1)

	fmt.Print("Masukkan koordinat dan radius lingkaran 2: ")
	fmt.Scanln(&pusatX2, &pusatY2, &radius2)

	fmt.Print("Masukkan koordinat titik: ")
	fmt.Scanln(&titikX, &titikY)

	// Mengecek posisi titik (titikX, titikY) terhadap lingkaran 1 dan lingkaran 2
	diLingkaran1 := cekdiDalamLingkaran(pusatX1, pusatY1, radius1, titikX, titikY)
	diLingkaran2 := cekdiDalamLingkaran(pusatX2, pusatY2, radius2, titikX, titikY)

	// Menentukan output berdasarkan posisi titik
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("\nTitik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("\nTitik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("\nTitik di dalam lingkaran 2")
	} else {
		fmt.Println("\nTitik di luar lingkaran 1 dan 2")
	}
}

```
### Output:
![Screenshot 2024-10-09 093408](https://github.com/user-attachments/assets/5b299ea7-d890-466e-bcd8-543da83725b1)

### Full code Screenshot:
![Screenshot 2024-10-09 093453](https://github.com/user-attachments/assets/98dd01ae-06f8-4621-9112-183ede33838c)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program

## Kesimpulan 

## Daftar Pustaka
[1] 


