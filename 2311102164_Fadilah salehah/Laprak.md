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
  Fadilah Salehah / 2311102164<br>
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
1. [Tujuan Praktikum](#tujuan-praktikum)
2. [Dasar Teori](#dasar-teori)
3. [Guided](#guided)
4. [Unguided](#unguided)
5. [Kesimpulan](#kesimpulan)
6. [Daftar Pustaka](#daftar-pustaka)

## Tujuan Praktikum
1. Mahasiswa diharapkan memahami pengertian fungsi dalam bahasa Go, termasuk struktur, sintaks, dan cara kerja fungsi.
2. Mahasiswa dapat mengimplementasikan fungsi dalam program Go untuk mengorganisir kode, meningkatkan keterbacaan, dan memisahkan logika program.
3. Mahasiswa dapat menjelaskan dan menerapkan penggunaan parameter dan return value dalam fungsi, serta memahami perbedaan antara pass-by-value dan pass-by-reference.
   
## Dasar Teori
**1. Dasar Fungsi dalam Bahasa Go**

- Dalam bahasa Go, fungsi merupakan first-class citizens, yang berarti mereka dapat disimpan dalam variabel, diteruskan sebagai argumen ke fungsi lain, dan bahkan dikembalikan dari fungsi.[1]

**2. Deklarasi Fungsi**
- Fungsi dideklarasikan menggunakan kata kunci `func`, diikuti oleh nama fungsi, parameter (jika ada), tipe kembalian (jika ada), dan blok kode yang akan dieksekusi.[2]

```go
func tambah(a int, b int) int {
    return a + b
}
```
- Fungsi `tambah` di atas menerima dua parameter `a` dan `b` bertipe `int`, kemudian mengembalikan hasil penjumlahan keduanya sebagai `int`.

**3. Multiple Return Values (Pengembalian Banyak Nilai)**
- Go mendukung pengembalian lebih dari satu nilai dari suatu fungsi. [2]
- Ini sering digunakan, terutama saat kita ingin mengembalikan hasil operasi bersama error.[2]

```go
func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("tidak bisa membagi dengan nol")
    }
    return a / b, nil
}
```
- Fungsi `bagi` mengembalikan dua nilai: hasil pembagian dan pesan error. 
- Jika terjadi pembagian dengan nol, fungsi mengembalikan error.

**4. Variadic Function (Fungsi dengan Argumen Tidak Terbatas)**
- Variadic function adalah fungsi yang dapat menerima jumlah argumen yang tidak terbatas. [1]
- Ini dicapai dengan menggunakan tiga titik `(...)` sebelum tipe parameter.[1]

```go
func jumlah(angka ...int) int {
    total := 0
    for _, n := range angka {
        total += n
    }
    return total
}
```

- Fungsi `jumlah` dapat menerima sejumlah argumen integer dan menjumlahkannya.

**5. Anonymous Function (Fungsi Anonim)**
- Fungsi anonim adalah fungsi tanpa nama yang bisa didefinisikan secara langsung, sering kali digunakan untuk tugas sederhana atau sebagai callback.[1]

```go
func main() {
    func() {
        fmt.Println("Ini adalah fungsi anonim")
    }()
}
```

- Fungsi anonim ini langsung dieksekusi setelah didefinisikan.

**6. Fungsi sebagai Argumen dan Nilai Kembalian**
- Fungsi dalam Go juga dapat diteruskan sebagai argumen ke fungsi lain atau dikembalikan sebagai nilai.[2]

```go
func apply(fn func(int) int, value int) int {
    return fn(value)
}

func kaliDua(x int) int {
    return x * 2
}

func main() {
    result := apply(kaliDua, 5)
    fmt.Println(result)  // Output: 10
}
```

- Fungsi `apply` menerima fungsi `kaliDua` sebagai argumen dan menerapkannya pada nilai `5`.

**7. Error Handling dengan Panic dan Recover**
- `panic` digunakan untuk menghentikan program, sedangkan `recover` untuk memulihkan eksekusi program dari kondisi panic.[1]

```go
func safeDivision(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    if b == 0 {
        panic("division by zero")
    }
    fmt.Println(a / b)
}

func main() {
    safeDivision(4, 2)  // Output: 2
    safeDivision(4, 0)  // Output: Recovered from: division by zero
}
```

- Fungsi `safeDivision` menggunakan `panic` untuk menghentikan eksekusi jika pembagian dengan nol terjadi, dan `recover` untuk menangani kondisi tersebut.


## Guided 

### 1. Program fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi.

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
![g1](https://github.com/user-attachments/assets/306f40c2-cba4-47b9-83c2-b793d949fdd0)

### Deskripsi Program : 
Program ini menerima dua bilangan bulat dari pengguna, lalu memeriksa bilangan mana yang lebih besar. 
Program akan menghitung nilai permutasi dari bilangan yang lebih besar sebagai 𝑛 dan bilangan yang lebih kecil sebagai 𝑟.
Hasil dari permutasi dihitung menggunakan rumus 𝑃(𝑛,𝑟)= 𝑛!/(𝑛−𝑟)!, dimana 𝑛! adalah faktorial dari 𝑛 dan (𝑛−𝑟)! adalah faktorial 𝑛−𝑟.


### 2. Program fungsi yang digunakan untuk menghitung luas dan keliling sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna

### Source Code :
```go
package main

import "fmt"

// Fungsi untuk menghitung luas persegi
func hitungLuas(sisi float64) float64 {
	return sisi * sisi
}

// Fungsi untuk menghitung keliling persegi
func hitungKeliling(sisi float64) float64 {
	return 4 * sisi
}

func main() {
	var sisi float64

	// Tanya pengguna untuk memasukkan sisi persegi
	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi)

	// Hitung dan tampilkan luas dan keliling
	luas := hitungLuas(sisi)
	keliling := hitungKeliling(sisi)

	fmt.Printf("Luas persegi adalah: %.2f\n", luas)
	fmt.Printf("Keliling persegi adalah: %.2f\n", keliling)
}
```

### Output:
![g2](https://github.com/user-attachments/assets/fc8ea075-3fa7-433e-b6b9-ee90cde8d17d)

### Deskripsi Program : 
Program ini menerima input berupa panjang sisi persegi dari pengguna, kemudian menghitung luas dan kelilingnya dengan menggunakan rumus :

- Luas Persegi : sisi x sisi
- Keliling Persegi : 4 x sisi
  
  
## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
**Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap
𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.**

**Catatan: permutasi (𝑃) dan kombinasi (𝐶) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!**

![image](https://github.com/user-attachments/assets/4978402f-272c-4bae-8ac4-e554843517a4)

Contoh : 

![image](https://github.com/user-attachments/assets/b7167807-34f3-4c61-91e9-3bdcc494c253)

### Source Code :
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("================================")
	fmt.Println("Program untuk Menghitung Permutasi & Kombinasi")
	fmt.Println("================================")

	// Mendapatkan input dari pengguna
	var a, b, c, d int
	a, b, c, d = ambilInput()

	// Validasi masukan
	if !validasiInput(a, b, c, d) {
		fmt.Println("Input tidak valid. Pastikan a >= c dan b >= d.")
		return
	}

	// Menampilkan hasil perhitungan
	tampilkanHasil(a, b, c, d)
}

// Fungsi untuk mengambil input dari pengguna
func ambilInput() (int, int, int, int) {
	var inputA, inputB, inputC, inputD string
	fmt.Print("Masukkan nilai a (bilangan asli): ")
	fmt.Scanln(&inputA)
	fmt.Print("Masukkan nilai b (bilangan asli): ")
	fmt.Scanln(&inputB)
	fmt.Print("Masukkan nilai c (bilangan asli): ")
	fmt.Scanln(&inputC)
	fmt.Print("Masukkan nilai d (bilangan asli): ")
	fmt.Scanln(&inputD)

	// Konversi input string menjadi integer
	a, _ := strconv.Atoi(inputA)
	b, _ := strconv.Atoi(inputB)
	c, _ := strconv.Atoi(inputC)
	d, _ := strconv.Atoi(inputD)

	return a, b, c, d
}

// Fungsi untuk memvalidasi apakah input sesuai syarat
func validasiInput(a, b, c, d int) bool {
	return a >= c && b >= d
}

// Fungsi untuk menghitung faktorial dari sebuah bilangan
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung nilai permutasi P(n, r)
func hitungPermutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung nilai kombinasi C(n, r)
func hitungKombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

// Fungsi untuk menampilkan hasil perhitungan permutasi dan kombinasi
func tampilkanHasil(a, b, c, d int) {
	fmt.Println("--------------------------------")
	fmt.Println("Hasil Perhitungan:")
	fmt.Printf("Permutasi P(%d, %d) = %d!/(%d-%d)! = %d\n", a, c, a, a, c, hitungPermutasi(a, c))
	fmt.Printf("Kombinasi C(%d, %d) = %d!/((%d!)*(%d-%d)!) = %d\n", a, c, a, c, a, c, hitungKombinasi(a, c))
	fmt.Printf("Permutasi P(%d, %d) = %d!/(%d-%d)! = %d\n", b, d, b, b, d, hitungPermutasi(b, d))
	fmt.Printf("Kombinasi C(%d, %d) = %d!/((%d!)*(%d-%d)!) = %d\n", b, d, b, d, b, d, hitungKombinasi(b, d))
	fmt.Println("--------------------------------")
}
```

### Output:
![un1](https://github.com/user-attachments/assets/b0aa1592-520e-4bf9-a1aa-ff4a4014a48e)

### Deskripsi Program : 
rogram ini adalah tentang menghitung permutasi dan kombinasi dari dua set bilangan yang dimasukkan oleh pengguna. Permutasi dan kombinasi adalah konsep matematika yang sering digunakan dalam teori peluang dan statistika.
Permutasi menghitung jumlah susunan berbeda dari sekumpulan elemen di mana urutan elemen tersebut penting.
Kombinasi menghitung jumlah cara memilih elemen dari suatu kumpulan, tanpa memperhatikan urutannya.

     
### 2. Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥2 , 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 + 1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔(ℎ (𝑥))).Tuliskan 𝑓 (𝑥), 𝑔 (𝑥), ℎ (𝑥) bentuk function.

**Masukan** terdiri dari sebuah bilangan bulat 𝑎 , 𝑏 dan 𝑐 yang dipisahkan oleh spasi.

**Keluaran** terdiri dari tiga baris. Baris pertama adalah (𝑓𝑜𝑔𝑜ℎ)(𝑎), baris kedua (𝑔𝑜ℎ𝑜𝑓)(𝑏), baris ketiga adalah (ℎ𝑜𝑓𝑜𝑔)(𝑐) !

Contoh :

![image](https://github.com/user-attachments/assets/b1baeb3b-31f8-4846-8b36-14c59e714cce)

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	// Menerima input nilai a dan menghitung f(g(h(a)))
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Println("Hasil (fogoh)(a) =", f(g(h(a))))

	// Menerima input nilai b dan menghitung g(h(f(b)))
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Println("Hasil (gohof)(b) =", g(h(f(b))))

	// Menerima input nilai c dan menghitung h(f(g(c)))
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)
	fmt.Println("Hasil (hofog)(c) =", h(f(g(c))))
}
```

### Output:


### Deskripsi Program : 
Program ini menggunakan konsep komposisi fungsi untuk menghitung hasil dari beberapa kombinasi fungsi yang berbeda berdasarkan input dari pengguna.
![un2](https://github.com/user-attachments/assets/de072548-f79b-41d0-b87d-32d094f8c147)

### 3. Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut.
**Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".**

Contoh :

![image](https://github.com/user-attachments/assets/5c0559d7-0487-43f3-8d6e-f9d6d095be5c)

![image](https://github.com/user-attachments/assets/a86ba834-47af-47ff-8953-281ee835f632)

### Source Code :
```go
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

// Fungsi untuk menentukan apakah titik (x, y) berada di dalam lingkaran
func diDalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1 int // Koordinat pusat dan radius lingkaran 1
	var cx2, cy2, r2 int // Koordinat pusat dan radius lingkaran 2
	var x, y int         // Koordinat titik yang akan dicek

	// Input data lingkaran 1
	fmt.Print("Masukkan koordinat pusat (cx1, cy1) dan radius r1 untuk lingkaran 1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input data lingkaran 2
	fmt.Print("Masukkan koordinat pusat (cx2, cy2) dan radius r2 untuk lingkaran 2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input titik yang akan diperiksa
	fmt.Print("Masukkan koordinat titik (x, y): ")
	fmt.Scan(&x, &y)

	// Memeriksa apakah titik berada di dalam lingkaran 1 atau lingkaran 2
	diDalamL1 := diDalamLingkaran(cx1, cy1, r1, x, y)
	diDalamL2 := diDalamLingkaran(cx2, cy2, r2, x, y)

	// Menampilkan hasil sesuai kondisi
	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output:
![un3](https://github.com/user-attachments/assets/cb7b0273-2050-4ebb-8e54-955c156e78f9)


### Deskripsi Program : 
Program ini digunakan untuk menentukan apakah suatu titik berada di dalam lingkaran, yang dapat digunakan dalam bidang geometri, grafik komputer, atau simulasi fisik.

## Kesimpulan 
Berdasarkan hasil praktikum fungsi pada bahasa pemrograman Go, dapat disimpulkan bahwa pemahaman tentang fungsi sangat penting dalam membangun program yang terstruktur dan efisien. 
Dari praktikum ini, dapat disimpulkan sebagai berikut:

1. Memahami konsep dasar fungsi yang memungkinkan pemisahan logika program menjadi bagian-bagian yang lebih kecil dan terorganisir.
2. Mengimplementasikan berbagai algoritma melalui fungsi, seperti perhitungan permutasi, kombinasi, dan geometri, untuk menyelesaikan masalah matematis secara sistematis.
3. Menerapkan interaksi pengguna melalui input dan output, sehingga program dapat menerima data dari pengguna dan menghasilkan hasil yang diharapkan.

## Daftar Pustaka
[1] A. A. A. Donovan and B. W. Kernighan, *The Go Programming Language*. Boston, MA: Addison-Wesley, 2015.

[2] W. Kennedy, B. Ketelsen, and E. St. Martin, *Go in Action*. Greenwich, CT: Manning Publications, 2016.
