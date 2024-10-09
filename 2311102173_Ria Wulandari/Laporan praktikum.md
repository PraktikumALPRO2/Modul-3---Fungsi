
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
  Ria Wulandari / 2311102173<br>
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
#### 1. Definisi Fungsi (Function)
Fungsi dalam pemrograman adalah blok kode yang dirancang untuk menjalankan tugas tertentu. Dalam Golang, fungsi merupakan first-class citizen, yang berarti fungsi dapat diperlakukan seperti variabel (dapat disimpan dalam variabel, dilewatkan sebagai argumen, atau dikembalikan oleh fungsi lain). Fungsi membantu memecah program menjadi bagian-bagian lebih kecil dan lebih mudah dipelihara, sehingga kode menjadi lebih modular, reusable, dan mudah dipahami.
#### 2. Deklarasi Fungsi
Deklarasi fungsi dalam Golang memiliki struktur sebagai berikut:
```go
func namaFungsi(parameter1 tipe, parameter2 tipe, ...) tipePengembalian {
    // Blok kode fungsi
    return nilai
}
```
1. `func`: Kata kunci yang digunakan untuk mendeklarasikan fungsi.
2. `namaFungsi`: Nama dari fungsi yang mengikuti konvensi penamaan variabel (huruf kecil untuk akses internal, huruf besar untuk akses luar paket).
3. `parameter`: Daftar parameter (opsional) yang diterima oleh fungsi, terdiri dari nama dan tipe.
4. `tipePengembalian`: Tipe data yang dikembalikan oleh fungsi (jika ada). Jika fungsi tidak mengembalikan nilai, ini bisa dihilangkan.
5. `Blok kode`: Berisi serangkaian instruksi yang akan dieksekusi saat fungsi dipanggil.
##### Contoh deklarasi fungsi
```go
func tambah(a int, b int) int {
    return a + b
}
```
Fungsi `tambah` menerima dua parameter `a` dan `b` dengan tipe data int, kemudian mengembalikan hasil penjumlahan dari keduanya.
#### 3. Pemanggilan Fungsi
Pemanggilan fungsi dilakukan dengan menggunakan nama fungsi diikuti dengan tanda kurung (), serta mengisi parameter yang diperlukan sesuai dengan deklarasi fungsi. Jika fungsi mengembalikan nilai, nilai tersebut bisa disimpan dalam variabel atau digunakan secara langsung.
##### Contoh pemanggilan fungsi
```go
hasil := tambah(3, 4)
fmt.Println(hasil) // Output: 7
```
## Guided
### 1. Membuat sebuah programbeserta fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi
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
#### Output


#### Deskripsi Program
Program ini dirancang untuk menghitung permutasi dari dua angka yang dimasukkan oleh pengguna. Ketika pengguna memasukkan dua angka, program akan memeriksa mana yang lebih besar. Jika angka pertama lebih besar atau sama dengan yang kedua, program akan menghitung permutasi menggunakan angka pertama terhadap angka kedua. Jika angka kedua lebih besar, program akan melakukan perhitungan permutasi sebaliknya, yaitu dari angka kedua terhadap angka pertama.

Untuk menjalankan perhitungan ini, program menggunakan dua fungsi. Fungsi pertama, `faktorial`, bertugas menghitung hasil faktorial dari sebuah angka. Fungsi kedua, `permutasi`, memanfaatkan hasil dari faktorial untuk menentukan permutasi. Setelah perhitungan selesai, program menampilkan hasilnya ke layar sehingga pengguna dapat melihat nilai permutasi yang dihasilkan. Program ini bekerja dengan cara yang mudah dan otomatis, membantu pengguna menghitung permutasi tanpa perlu memahami proses matematis yang terlibat.
#### Algoritma Program
1. Mulai program.
2. Deklarasikan dua variabel untuk input pengguna, `a` dan `b`.
3. Minta pengguna untuk memasukkan dua angka bulat.
4. Jika `a` lebih besar atau sama dengan `b`, hitung permutasi dengan `a` dan `b`.
5. Jika tidak, hitung permutasi dengan `b` dan `a`.
6. Untuk menghitung permutasi, panggil fungsi yang menghitung faktorial dari `n`.
7. Kembalikan hasil faktorial untuk menghitung permutasi.
8. Tampilkan hasil permutasi kepada pengguna.
9. Program selesai.
#### Cara Kerja Program
1. Program dimulai dan dua variabel, `a` dan `b`, dideklarasikan untuk menyimpan input dari pengguna.
2. Pengguna diminta untuk memasukkan dua angka bulat yang akan digunakan dalam perhitungan.
3. Program kemudian memeriksa mana dari kedua angka tersebut yang lebih besar.
4. Jika angka `a` lebih besar atau sama dengan `b`, program akan menghitung permutasi menggunakan `a` sebagai `n` dan `b` sebagai `r`.
5. Jika `b` lebih besar dari `a`, program akan menghitung permutasi dengan `b` sebagai `n` dan `a` sebagai r.
6. Untuk menghitung permutasi, program memanggil fungsi `faktorial`, yang menghitung hasil kali semua angka dari 1 hingga n.
7. Setelah menghitung faktorial dari n dan (n-r), hasilnya digunakan untuk menghitung permutasi.
8. Hasil permutasi kemudian ditampilkan kepada pengguna.
9. Program berakhir setelah menampilkan hasil.
### 2. Program fungsi yang digunakan untuk menghitung luas dan keliling sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna
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
#### Output


#### Deskripsi Program
Program ini menghitung luas dan keliling sebuah persegi. Pertama, program meminta pengguna untuk memasukkan panjang sisi persegi. Setelah menerima input, program menghitung luas dengan mengalikan panjang sisi dengan dirinya sendiri dan menghitung keliling dengan mengalikan panjang sisi dengan empat. Akhirnya, program menampilkan hasil perhitungan luas dan keliling kepada pengguna. Program ini sangat berguna untuk memahami cara menghitung dua sifat dasar dari bentuk geometri persegi.
#### Algoritma Program
1. Mulai program dengan mendeklarasikan fungsi utama `main()` sebagai titik masuk eksekusi program.
2. Deklarasikan variabel `panjangSisi` dengan tipe `float64`, untuk menyimpan nilai panjang sisi persegi yang akan dimasukkan oleh pengguna.
3. Tampilkan pesan ke pengguna dengan tulisan: "Masukkan panjang sisi persegi: ", untuk meminta input dari pengguna.
4. Gunakan `fmt.Scan()` untuk membaca input dari pengguna dan simpan nilai tersebut ke dalam variabel `panjangSisi`.
5. Deklarasikan variabel `LuasPersegi` dan `KelilingPersegi`, masing-masing dengan tipe `float64`, untuk menyimpan hasil perhitungan luas dan keliling persegi.
6. Hitung luas persegi dengan mengalikan `panjangSisi` dengan dirinya sendiri, dan simpan hasilnya dalam variabel `LuasPersegi`.
7. Hitung keliling persegi dengan mengalikan `panjangSisi` dengan 4, kemudian simpan hasilnya dalam variabel `KelilingPersegi`.
8. Tampilkan hasil perhitungan luas persegi kepada pengguna menggunakan `fmt.Printf()` dengan format yang jelas.
9 Tampilkan hasil perhitungan keliling persegi kepada pengguna menggunakan `fmt.Printf()`, agar pengguna dapat melihat kedua hasil perhitungan tersebut.
10. Akhiri program setelah semua hasil ditampilkan dengan sukses.
#### Cara Kerja Program
1. Program dimulai dengan mendeklarasikan fungsi utama `main()`, yang menjadi titik awal eksekusi program.

2. Sebuah variabel bernama `panjangSisi` dideklarasikan dengan tipe data `float64` untuk menyimpan input panjang sisi persegi dari pengguna.

3. Program meminta pengguna untuk memasukkan panjang sisi persegi melalui pesan "Masukkan panjang sisi persegi: ".

4. Fungsi `fmt.Scan()` digunakan untuk membaca input dari pengguna dan menyimpannya dalam variabel `panjangSisi`.

5. Luas persegi dihitung dengan mengalikan nilai `panjangSisi` dengan dirinya sendiri, dan hasilnya disimpan dalam variabel `LuasPersegi`.

6. Keliling persegi dihitung dengan mengalikan nilai `panjangSisi` dengan 4, dan hasilnya disimpan dalam variabel `KelilingPersegi`.

7. Program menampilkan hasil perhitungan luas persegi kepada pengguna menggunakan fungsi `fmt.Printf()` dengan format yang sesuai.

8. Hasil perhitungan keliling persegi juga ditampilkan kepada pengguna menggunakan fungsi `fmt.Printf()` dengan format yang serupa.

9. Program selesai dieksekusi setelah semua hasil perhitungan ditampilkan kepada pengguna, tanpa memerlukan input tambahan.

## Unguided
### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? 
Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap 𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.

Catatan: permutasi (𝑃) dan kombinasi (𝐶) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!
![Cuplikan layar 2024-10-09 092504](https://github.com/user-attachments/assets/04069b4d-c3a6-403a-8f92-735d834a7605)
![Cuplikan layar 2024-10-09 092732](https://github.com/user-attachments/assets/1343ac89-a62b-4b2d-91f9-0c7446c6ab66)

#### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi P(n, r)
func permutation(n, r int) int {
	if r > n {
		return 0
	}
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r)
func combination(n, r int) int {
	if r > n {
		return 0
	}
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	// Input
	var a, b, c, d int
	fmt.Print("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	// Baris pertama: permutasi dan kombinasi dari a dan c
	Pac := permutation(a, c)
	Cac := combination(a, c)

	// Baris kedua: permutasi dan kombinasi dari b dan d
	Pbd := permutation(b, d)
	Cbd := combination(b, d)

	// Output
	fmt.Println(Pac, Cac)
	fmt.Println(Pbd, Cbd)
}
```
#### Output


#### Deskripsi Program
Program ini dirancang untuk menghitung permutasi dan kombinasi dari empat angka yang dimasukkan oleh pengguna. Pertama, program memiliki fungsi `factorial` untuk menghitung faktorial dari sebuah angka. Fungsi `permutation` menghitung jumlah cara untuk mengatur objek, sementara fungsi `combination` menghitung jumlah cara untuk memilih objek tanpa memperhatikan urutan. Setelah pengguna memasukkan empat angka, program menghitung dan menampilkan hasil permutasi dan kombinasi untuk dua pasangan angka tersebut dalam dua baris terpisah, sehingga memudahkan pengguna memahami konsep ini.
#### Algoritma Program
1. Mulai program dan deklarasikan variabel `a`, `b`, `c`, dan `d`.

2. Definisikan fungsi `factorial` untuk menghitung faktorial dari `n`.

3. Definisikan fungsi `permutation` untuk menghitung permutasi dari `n` dan `r`.

4. Definisikan fungsi `combination` untuk menghitung kombinasi dari `n` dan `r`.

5. Dalam fungsi `main`, minta input dari pengguna untuk `a`, `b`, `c`, dan `d`.

6. Hitung permutasi dan kombinasi untuk pasangan `(a, c)` dan `(b, d)`.

7. Tampilkan hasil dari permutasi dan kombinasi yang telah dihitung.

8. Selesai program.
#### Cara Kerja Program
1. Program dimulai dengan mendeklarasikan variabel untuk menyimpan input nilai `a`, `b`, `c`, dan `d`.

2. Fungsi `factorial` didefinisikan untuk menghitung faktorial dari bilangan yang diberikan. Fungsi ini mengembalikan `1` jika nilai yang dimasukkan adalah `0` atau `1`, dan menghitung faktorial dengan menggunakan loop untuk nilai lebih besar dari `1``.

3. Fungsi permutation` didefinisikan untuk menghitung permutasi, di mana fungsi ini akan memanggil fungsi `factorial` untuk menghitung nilai faktorial `n` dan `n-r`, kemudian membagi hasilnya untuk mendapatkan nilai permutasi.

4. Fungsi `combination` didefinisikan untuk menghitung kombinasi. Fungsi ini juga memanggil fungsi `factorial` untuk menghitung faktorial dari `n`, `r`, dan `n-r`, kemudian membagi hasilnya sesuai dengan rumus kombinasi.

5. Dalam fungsi `main`, program meminta input dari pengguna untuk nilai `a`, `b`, `c`, dan `d`.

6. Program menghitung permutasi dan kombinasi untuk pasangan `(a, c)` dan `(b, d)` menggunakan fungsi yang telah didefinisikan sebelumnya.

7. Hasil dari permutasi dan kombinasi yang telah dihitung ditampilkan ke layar.

8. Program selesai setelah menampilkan hasil.
### 2. Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥2 , 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 + 1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔(ℎ (𝑥))).Tuliskan 𝑓 (𝑥), 𝑔 (𝑥), ℎ (𝑥) bentuk function.
Masukan terdiri dari sebuah bilangan bulat 𝑎 , 𝑏 dan 𝑐 yang dipisahkan oleh spasi.

Keluaran terdiri dari tiga baris. Baris pertama adalah (𝑓𝑜𝑔𝑜ℎ)(𝑎), baris kedua (𝑔𝑜ℎ𝑜𝑓)(𝑏), baris ketiga adalah (ℎ𝑜𝑓𝑜𝑔)(𝑐) !
![image](https://github.com/user-attachments/assets/96f47f07-fecb-44e4-b833-d665e9f8f414)
#### Source Code :
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

// Fungsi komposisi fogoh = f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi komposisi gohof = g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi komposisi hofog = h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// Output hasil komposisi fungsi
	fmt.Println(fogoh(a))  // Baris pertama (fogoh(a))
	fmt.Println(gohof(b))  // Baris kedua (gohof(b))
	fmt.Println(hofog(c))  // Baris ketiga (hofog(c))
}
```
#### Output


#### Deskripsi Program
Program ini ditulis dalam bahasa Go untuk menghitung hasil dari beberapa komposisi fungsi matematika. Terdapat tiga fungsi dasar: `f(x)` yang mengkuadratkan `x`, `g(x)` yang mengurangi `x` dengan dua, dan `h(x)` yang menambahkan satu pada `x`. Program meminta pengguna untuk memasukkan tiga bilangan bulat, kemudian menghitung dan menampilkan hasil dari komposisi fungsi `f(g(h(x)))`, `g(h(f(x)))`, dan `h(f(g(x)))` untuk input tersebut. Hasil dari setiap komposisi ditampilkan secara terpisah.
#### Algoritma Program
