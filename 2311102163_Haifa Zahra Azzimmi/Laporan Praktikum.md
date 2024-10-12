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
**1. Definsi Fungsi**

   Definisi fungsi dalam mata kuliah algoritma pemrograman adalah sekumpulan instruksi yang dikelompokkan bersama untuk melakukan tugas tertentu dan dapat dipanggil dari berbagai bagian 
   dalam program. Fungsi membantu dalam pengorganisasian kode, membuatnya lebih modular dan dapat digunakan kembali, sehingga memudahkan pemeliharaan dan pengembangan program.

**2. Penerapan Fungsi**

   Penerapan fungsi dalam pemrograman adalah tentang bagaimana kita menggunakan fungsi-fungsi untuk menyederhanakan, mengorganisir, dan memecah program besar menjadi bagian-bagian yang 
   lebih kecil dan dapat dikelola.
   
**3. Deklarasi Fungsi**

   Deklarasi fungsi adalah proses dalam pemrograman di mana kita mendefinisikan nama fungsi, tipe parameter yang akan diterima, dan tipe data yang akan dikembalikan. Deklarasi ini 
   penting untuk memberi tahu kompiler tentang keberadaan fungsi tersebut dan bagaimana fungsi itu akan berinteraksi dengan bagian lain dari program.
   
   Contoh deklarasi fungsi dalam bahasa Go adalah sebagai berikut:
   ```go
   // Deklarasi fungsi untuk menghitung faktorial
	func faktorial(n int) int {
    // Implementasi fungsi di sini
	}
   ```
   Dalam deklarasi ini, faktorial adalah nama fungsi, n int menunjukkan bahwa fungsi ini menerima satu parameter bernama n dengan tipe data integer, dan int setelah tanda kurung 
   menandakan bahwa fungsi ini akan mengembalikan nilai bertipe integer.
   
**4. Pemanggilan Fungsi**

   Pemanggilan fungsi adalah cara kita menggunakan atau menjalankan fungsi yang telah kita deklarasikan sebelumnya di dalam program kita. Ini dilakukan dengan menuliskan nama fungsi dan 
   memberikan argumen yang sesuai jika diperlukan.

   Contohnya, jika kita memiliki fungsi faktorial yang kita deklarasikan sebelumnya, kita dapat memanggilnya seperti ini:
   ```go
   // Memanggil fungsi faktorial dengan argumen 5
   hasil := faktorial(5)
   fmt.Println("Hasil faktorial dari 5 adalah:", hasil)
   ```

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
5. 

## Unguided
### 1. Program untuk menghitung Faktorial, Permutasi dan Kombinasi
![image](https://github.com/user-attachments/assets/c4cbd2b7-d81e-430d-9b8a-a85ad0ea5fb8)
### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
	if n < r {
		return 0 // Tidak valid
	}
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int) int {
	if n < r {
		return 0 // Tidak valid
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan nilai a dan b (dengan a >= c dan b >= d): ")
	fmt.Scan(&a, &b)
	fmt.Print("Masukkan nilai c dan d: ")
	fmt.Scan(&c, &d)

	if a >= c && b >= d {
		// Baris pertama: Hasil permutasi dan kombinasi a terhadap c
		fmt.Printf("Permutasi P(%d, %d) = %d\n", a, c, permutasi(a, c))
		fmt.Printf("Kombinasi C(%d, %d) = %d\n", a, c, kombinasi(a, c))

		// Baris kedua: Hasil permutasi dan kombinasi b terhadap d
		fmt.Printf("Permutasi P(%d, %d) = %d\n", b, d, permutasi(b, d))
		fmt.Printf("Kombinasi C(%d, %d) = %d\n", b, d, kombinasi(b, d))
	} else {
		fmt.Println("Pastikan bahwa a >= c dan b >= d.")
	}
}

```

### Output :
![image](https://github.com/user-attachments/assets/3fe41467-2bc5-40b9-8464-c79560f2ab48)
![image](https://github.com/user-attachments/assets/c8532484-4d5c-4e60-b736-bb244ec38c2c)

### Full code Screenshot :
![image](https://github.com/user-attachments/assets/155aec83-363f-49fc-b5c9-1e5c18b20f96)

#### Deskripsi Program	:
Program ini meminta pengguna untuk memasukkan nilai a, b, c, dan d. Setelah itu, program menghitung dan menampilkan hasil permutasi dan kombinasi dari a terhadap c, serta b terhadap d, dengan menggunakan fungsi faktorial, permutasi, dan kombinasi. Jika a tidak lebih besar atau sama dengan c, atau b tidak lebih besar atau sama dengan d, program akan memberikan pesan error dan meminta pengguna untuk memastikan bahwa a >= c dan b >= d.

### Algoritma Program :
1. Mulai
2. Meminta pengguna untuk memasukkan nilai a dan b.
3. Meminta pengguna untuk memasukkan nilai c dan d.
4. Memeriksa apakah a lebih besar atau sama dengan c dan b lebih besar atau sama dengan d.
5. Menghitung permutasi a terhadap c menggunakan fungsi permutasi.
6. Menghitung kombinasi a terhadap c menggunakan fungsi kombinasi.
7. Menampilkan hasil permutasi dan kombinasi a terhadap c.
8. Menghitung permutasi b terhadap d menggunakan fungsi permutasi.
9. Menghitung kombinasi b terhadap d menggunakan fungsi kombinasi.
10. Menampilkan hasil permutasi dan kombinasi b terhadap d.

### Cara Kerja Program :
1. Algoritma dimulai dengan mendefinisikan fungsi-fungsi yang diperlukan untuk menghitung faktorial, permutasi, dan kombinasi.
2. Di dalam fungsi main, program meminta input dari pengguna untuk dua pasang angka.
3. Setelah memvalidasi bahwa input memenuhi syarat, program menghitung dan menampilkan hasil permutasi dan kombinasi.
4. Jika syarat tidak terpenuhi, program akan memberikan pesan peringatan kepada pengguna.


### 2. 
![image](https://github.com/user-attachments/assets/afb8ecf4-9901-4764-90ef-56658510ffa3)

### Source Code :
```go
package main

import (
	"fmt"
)

// Definisikan fungsi f(x), g(x), h(x)
func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

// Fungsi komposisi untuk fogoh(x), gohof(x), dan hofog(x)
func fogoh(x int) int {
	return f(g(h(x)))
}

func gohof(x int) int {
	return g(h(f(x)))
}

func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int

	// Input bilangan bulat a, b, c
	fmt.Print("Masukkan nilai a, b, c (dipisahkan dengan spasi): ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	// Output dengan keterangan
	fmt.Printf("Hasil fogoh(%d) = %d\n", a, fogoh(a))    // fogoh(a)
	fmt.Printf("Hasil gohof(%d) = %d\n", b, gohof(b))    // gohof(b)
	fmt.Printf("Hasil hofog(%d) = %d\n", c, hofog(c))    // hofog(c)
}
```
### Output :
![image](https://github.com/user-attachments/assets/40e06e72-264d-4dfe-bc7e-55c08d3d95b3)

### Full code Screenshot :
![image](https://github.com/user-attachments/assets/24aa283f-4b91-4301-8a4a-8f754ac5cbd8)

### Deskripsi Program :
Program ini mendefinisikan tiga fungsi matematika: f(x), g(x), dan h(x). Fungsi-fungsi ini digunakan untuk menghitung nilai dari komposisi fungsi fogoh(x), gohof(x), dan hofog(x). Program ini meminta pengguna untuk memasukkan tiga nilai integer, yaitu a, b, dan c, dan kemudian menghitung serta menampilkan hasil dari ketiga komposisi fungsi tersebut untuk masing-masing nilai yang dimasukkan.

### Algoritma Program :
1. Mulai
2. Definisikan fungsi
   - `f(x)`: \( f(x) = x^2 \)
   - `g(x)`: \( g(x) = x - 2 \)
   - `h(x)`: \( h(x) = x + 1 \)
3. Definisikan fungsi komposisi
   - `fogoh(x)`: \( f(g(h(x))) \)
   - `gohof(x)`: \( g(h(f(x))) \)
   - `hofog(x)`: \( h(f(g(x))) \)
4. Input: Minta pengguna memasukkan nilai `a`, `b`, `c`.
5. Proses
   - Hitung dan simpan hasil:
     - `hasil1 = fogoh(a)`
     - `hasil2 = gohof(b)`
     - `hasil3 = hofog(c)`
6. Output Tampilkan `hasil1`, `hasil2`, dan `hasil3`.
7. Selesai
    
### Cara Kerja Program :
Program ini bekerja dengan meminta pengguna untuk memasukkan tiga nilai integer: a, b, dan c. Setelah itu, program menghitung dan menampilkan hasil dari tiga komposisi fungsi matematika: fogoh(x), gohof(x), dan hofog(x).

Untuk setiap nilai yang dimasukkan, program:

1. Menghitung nilai fogoh(a), yang merupakan komposisi fungsi dari f(g(h(a))).

2. Menghitung nilai gohof(b), yang merupakan komposisi fungsi dari g(h(f(b))).

3. Menghitung nilai hofog(c), yang merupakan komposisi fungsi dari h(f(g(c))).

Hasil perhitungan ini kemudian ditampilkan di layar dengan deskripsi yang sesuai. Jadi intinya: program minta input, menghitung komposisi fungsi, dan menampilkan hasilnya.


### 3. 
![image](https://github.com/user-attachments/assets/1a98819f-2087-4af4-9d6f-986e9b30323e)

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
		fmt.Println("\nTitik tersebut berada di dalam kedua lingkaran")
	} else if diLingkaran1 {
		fmt.Println("\nTitik tersebut berada di dalam lingkaran pertama saja")
	} else if diLingkaran2 {
		fmt.Println("\nTitik tersebut berada di dalam lingkaran kedua saja")
	} else {
		fmt.Println("\nTitik tersebut tidak berada di dalam kedua lingkaran")
	}
}


```
### Output :
![image](https://github.com/user-attachments/assets/4058dd66-7afd-4086-9e38-130393fd08c1)

### Full code Screenshot :
![image](https://github.com/user-attachments/assets/6e412987-98c1-46e1-8970-b8668e7b2363)

### Deskripsi Program :
Program ini bertujuan untuk memeriksa apakah suatu titik berada di dalam satu atau dua lingkaran yang diberikan. Program menerima koordinat dan radius dari dua lingkaran serta koordinat dari sebuah titik, kemudian menentukan apakah titik tersebut berada di dalam salah satu, kedua, atau tidak berada di dalam lingkaran manapun.

### Algoritma Program :
1. Mulai:
   
   - Definisikan fungsi hitungJarak untuk menghitung jarak antara dua titik.
   - Definisikan fungsi cekdiDalamLingkaran untuk memeriksa apakah titik berada di dalam lingkaran berdasarkan jarak.

3. Input Pengguna:
   
   - Minta pengguna memasukkan koordinat dan radius untuk dua lingkaran.
   - Minta pengguna memasukkan koordinat titik yang akan diperiksa.

4. Periksa Posisi Titik:
   
   - Hitung apakah titik berada di dalam lingkaran pertama.
   - Hitung apakah titik berada di dalam lingkaran kedua.

5. Output Hasil:
   
   - Jika titik berada di dalam kedua lingkaran, tampilkan pesan "Titik berada di dalam kedua lingkaran".
   - Jika titik berada di dalam lingkaran pertama saja, tampilkan pesan "Titik berada di dalam lingkaran pertama saja".
   - Jika titik berada di dalam lingkaran kedua saja, tampilkan pesan "Titik berada di dalam lingkaran kedua saja".
   - Jika titik tidak berada di dalam kedua lingkaran, tampilkan pesan "Titik tidak berada di dalam kedua lingkaran".

5. Selesai.

### Cara Kerja Program :
1. Input Nilai: Pengguna memasukkan koordinat dan radius untuk dua lingkaran, serta koordinat untuk satu titik.

2. Hitung Jarak: Program menghitung jarak titik ke pusat lingkaran menggunakan rumus jarak Euclidean.

3. Cek Posisi Titik: Program memeriksa apakah titik tersebut berada di dalam salah satu atau kedua lingkaran berdasarkan jarak yang dihitung.

4. Output Hasil: Program menampilkan pesan yang menunjukkan apakah titik berada di dalam satu, kedua, atau tidak ada di dalam lingkaran manapun.

