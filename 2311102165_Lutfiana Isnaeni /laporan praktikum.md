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

### 1. Pengertian Fungsi
Fungsi dalam Go adalah blok kode yang dapat digunakan kembali untuk melakukan tugas tertentu. Fungsi mempermudah pembagian kode menjadi modul-modul yang lebih kecil, sehingga meningkatkan keterbacaan dan efisiensi.
### 2. Struktur Fungsi di Go
Fungsi di Go dideklarasikan menggunakan kata kunci `func`, diikuti oleh nama fungsi, parameter, dan tipe kembalian (jika ada). Contoh fungsi sederhana:
```go
func tambah(a int, b int) int {
    return a + b
}
```
Fungsi ini menerima dua parameter `a` dan `b`, keduanya bertipe `int`, dan mengembalikan hasil penjumlahan sebagai `int`.
### 3. Fungsi dengan Banyak Nilai Kembalian
Go mendukung fungsi yang bisa mengembalikan lebih dari satu nilai, sangat berguna dalam banyak kasus. Contoh:
```go
func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("pembagian dengan nol")
    }
    return a / b, nil
}
```

### 4. Fungsi Anonim
Fungsi dalam Go juga dapat dibuat tanpa nama (anonim) dan disimpan dalam variabel atau langsung dieksekusi. Contoh fungsi anonim:
```go
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(3, 4)) // Output: 7
```

### 5. Manfaat Fungsi dalam 
- Modularitas dan Keterbacaan: Mempermudah pembagian kode menjadi bagian-bagian yang lebih terstruktur.
- Penggunaan Ulang: Fungsi dapat dipanggil berkali-kali tanpa harus menulis ulang logika yang sama.
- Dukungan Nilai Kembalian Ganda: Memungkinkan pengembalian hasil dan pesan error dalam satu fungsi.

## Guided

### Soal Nomor 1.
	Membuat sebuah programbeserta fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi
#### Source Code

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

### Deskripsi Program:
	Kode di atas digumakan untuk menghitung nilai permutasi dari dua angka yang diinputkan oleh pengguna.
 	Permutasi menghitung berapa banyak cara mengatur objek dari n objek yang berbeda, dengan mempertimbangkan urutannya.

### Cara Kerja Program:

	1. Fungsi main():
	  - Program dimulai dengan membaca dua bilangan bulat, a dan b, dari input pengguna.
	  - Jika a lebih besar atau sama dengan b, program akan menghitung permutasi P(a, b).
	  - Jika sebaliknya, maka program menghitung permutasi P(b, a).
 	  - Hasil permutasi kemudian ditampilkan.

	2. Fungsi faktorial(n int) int:
	  - Fungsi ini digunakan untuk menghitung faktorial dari bilangan n.
	  - Faktorial adalah hasil perkalian dari bilangan bulat positif dari 1 hingga n.
   	    Algoritma yang digunakan adalah iterasi, di mana program melakukan perkalian bertahap dari 1 sampai n.
 	  - Misalnya, faktorial dari 5 (5!) dihitung sebagai 1 x 2 x 3 x 4 x 5, yang hasilnya adalah 120.

	3. Fungsi permutasi(n, r int) int:
 	  - Fungsi ini digunakan untuk menghitung permutasi P(n, r).
	  - Rumus permutasi adalah: P(n, r) = n! / (n-r)!, yang berarti menghitung faktorial dari n kemudian membaginya dengan faktorial dari n - r.
 	  - Fungsi permutasi() memanggil fungsi faktorial() untuk melakukan perhitungan faktorial dari n dan n-r.

### Algoritma Program:
 	 - Input: Dua bilangan bulat a dan b dimasukkan oleh pengguna.
 	 - Proses: Program membandingkan nilai a dan b untuk menentukan urutan input bagi fungsi permutasi. 
   	   Program menghitung permutasi menggunakan fungsi faktorial dengan iterasi.
 	 - Output: Hasil dari perhitungan permutasi ditampilkan sebagai output.

### Soal Nomor 2.
	Program fungsi yang digunakan untuk menghitung luas dan keliling sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna
#### Source Code
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

### Deskripsi Program:
	Program di atas adalah sebuah program yang dibuat untuk menghitung luas dan keliling persegi berdasarkan input panjang sisi dari pengguna.
 	Program ini mengikuti algoritma yang sederhana, yaitu menerima input, melakukan perhitungan luas dan keliling, lalu menampilkan hasilnya.

	Program ini bertujuan untuk menghitung dua hal dari sebuah persegi:

	1. Luas Persegi: Dihitung dengan rumus Luas=sisi × sisi
	2. Keliling Persegi: Dihitung dengan rumus Keliling= 4 × sisi
	   Program menggunakan input dari pengguna berupa panjang sisi persegi, lalu menghitung luas dan kelilingnya, dan menampilkan hasilnya ke layar.

### Algoritma Program:
	1. Program meminta pengguna untuk memasukkan panjang sisi dari persegi.
	2. Panjang sisi yang dimasukkan disimpan dalam variabel `panjangSisi.`
	3. Program kemudian menghitung luas persegi menggunakan rumus Luas = sisi × sisi dan menyimpannya dalam variabel LuasPersegi.
	4. Selanjutnya, program menghitung keliling persegi dengan rumus Keliling = 4 × sisi dan menyimpannya dalam variabel KelilingPersegi.
	5. Setelah menghitung, hasil perhitungan luas dan keliling ditampilkan menggunakan fungsi fmt.Printf.
 ### Cara Kerja Program:
	1. Input: Program meminta input dari pengguna, yaitu panjang sisi persegi. Misalnya, pengguna memasukkan 4.
	2. Proses: Program menghitung luas persegi dan keliling persegi berdasarkan rumus yang sesuai:
	   • Luas persegi = 7 × 7 = 49
	   • Keliling persegi = 4 × 7 = 28
	3. Output: Hasil perhitungan ditampilkan ke layar dalam format sebagai berikut:
	   • Luas persegi: 49
	   • Keliling persegi: 28

## UNGUIDED

### Soal Nomor 1. 
	Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi.
 	Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? 
	Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑.
	Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap 𝑐, 
	sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.
Catatan: permutasi (𝑃) dan kombinasi (𝐶) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!
![ss soal unguided 1 modul 3](https://github.com/user-attachments/assets/1ac36a6b-4e1b-4916-9f3d-f42d762e3f1a)

#### Source Code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

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
![ss output unguided 1 modul 3](https://github.com/user-attachments/assets/8cf8d7c4-a418-492d-80d0-2b2661c4401c)

### Deskripsi Program:
	Program di atas dibuat untuk menghitung permutasi dan kombinasi dari dua pasang nilai input yang diberikan oleh pengguna.
 	Program menggunakan rumus permutasi dan kombinasi dasar dari matematika. 
 	Program ini menerima empat input: a, b, c, dan d, lalu menghitung:
	1. Permutasi P (n, r) P (n, r) P (n, r): menghitung permutasi dari a dengan c, dan b dengan d.
	2. Kombinasi C (n, r) C (n, r) C(n, r): menghitung kombinasi dari a dengan c, dan b dengan d.
	Rumus yang Digunakan:
	1. Faktorial n! dihitung sebagai hasil perkalian dari semua bilangan positif dari 1 hingga n.
	   n!= 1 × 2 × 3 × ⋯ × n
	2. Permutasi P (n, r) P (n, r) P ( n, r): menghitung berapa cara mengatur r objek dari n objek yang berbeda, tanpa pengulangan.
	   P (n, r)= n!/(n − r)!
	3. Kombinasi C (n, r) C (n, r) C(n, r): menghitung berapa cara memilih r objek dari n objek yang berbeda, tanpa memperhatikan urutan.
	   C (n, r)= n! / r! × (n−r)!
    
### Algoritma Program:
	1. Program mendefinisikan fungsi factorial(n) untuk menghitung faktorial dari bilangan nnn.
	2. Fungsi permutation (n, r) digunakan untuk menghitung permutasi berdasarkan rumus permutasi.
	3. Fungsi combination (n, r) digunakan untuk menghitung kombinasi berdasarkan rumus kombinasi.
	4. Program meminta pengguna memasukkan empat nilai: a, b, c, dan d.
	5. Program menghitung permutasi dan kombinasi dari dua pasang nilai:
	   • Permutasi dan kombinasi dari a dan c
	   • Permutasi dan kombinasi dari b dan d
	6. Hasil perhitungan kemudian ditampilkan dalam dua baris output.

### Cara Kerja Program:
	
	1. Input: Pengguna diminta untuk memasukkan nilai a, b, c, dan d. Misalnya, pengguna memasukkan nilai a= 5, b= 10,  c= 3, d= 10 
	2. Proses:
	   • Menghitung permutasi P(5, 3) dan kombinasi C(5, 3).  
	   • Menghitung permutasi P(10, 10) dan kombinasi C(10, 10).
	3. Output: Program mencetak dua baris hasil perhitungan:
	   • Baris pertama menampilkan hasil permutasi dan kombinasi untuk a dan c.
	   •Baris kedua menampilkan hasil permutasi dan kombinasi untuk b dan d.
    
### Soal Nomor 2.
	Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥2 , 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 + 1.
	Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔(ℎ (𝑥))).Tuliskan 𝑓 (𝑥), 𝑔 (𝑥), ℎ (𝑥) bentuk function.
	Masukan terdiri dari sebuah bilangan bulat 𝑎 , 𝑏 dan 𝑐 yang dipisahkan oleh spasi.
	Keluaran terdiri dari tiga baris. Baris pertama adalah (𝑓𝑜𝑔𝑜ℎ)(𝑎), baris kedua (𝑔𝑜ℎ𝑜𝑓)(𝑏), baris ketiga adalah (ℎ𝑜𝑓𝑜𝑔)(𝑐) !
![ss soal unguided 2 modul 3](https://github.com/user-attachments/assets/b57e11bf-e828-4ed0-a251-e884d549d5a9)
#### Source Code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165
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

// Fungsi untuk menghitung komposisi f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi untuk menghitung komposisi g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi untuk menghitung komposisi h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	// Input tiga nilai a, b, c
	var a, b, c int
	fmt.Print("Masukkan nilai a, b, c: ")
	fmt.Scan(&a, &b, &c)

	// Baris pertama: f(g(h(a)))
	fmt.Println(fogoh(a))

	// Baris kedua: g(h(f(b)))
	fmt.Println(gohof(b))

	// Baris ketiga: h(f(g(c)))
	fmt.Println(hofog(c))
}
```
#### Output
![image](https://github.com/user-attachments/assets/acae0b73-b78b-46dc-8e2c-2f8aa4f7ea0b)

### Deskripsi Program:
	Program di atas bertujuan untuk menghitung komposisi beberapa fungsi matematika sederhana, yaitu fungsi `f(x)`, `g(x)`, dan `h(x)`. 
 	Program ini meminta pengguna untuk memasukkan tiga bilangan bulat `a`, `b`, dan `c`, kemudian menghitung hasil dari komposisi fungsi berikut:
	1. f(g(h(a)))
	2. g(h(f(b))) 
	3. h(f(g(c)))
	Fungsi-fungsi yang digunakan dalam program adalah:
	- **f(x)**: menghitung f(x) = x^2 
	- **g(x)**: menghitung g(x) = x - 2 
	- **h(x)**: menghitung h(x) = x + 1
### Algoritma Program:
	1. Deklarasi tiga fungsi matematika: `f(x)`, `g(x)`, dan `h(x)`.
 	  - ```f(x)``` menghitung kuadrat dari nilai `x`.
 	  - `g(x)` mengurangi 2 dari nilai `x`.
 	  - `h(x)` menambahkan 1 ke nilai `x`.
	2. Terdapat tiga fungsi komposisi yang memanggil fungsi-fungsi di atas:
	   - `fogoh(x)` menghitung komposisi f(g(h(x))).
	   - `gohof(x)` menghitung komposisi g(h(f(x))).
 	   - `hofog(x)` menghitung komposisi h(f(g(x))).
	3. Pada fungsi `main`, program meminta input tiga nilai `a`, `b`, dan `c` dari pengguna.
	4. Program kemudian menampilkan hasil dari komposisi fungsi-fungsi:
 	  - `fogoh (a)` menghitung f(g(h(a))) dan menampilkan hasilnya.
  	  - `gohof (b)` menghitung g(h(f(b))) dan menampilkan hasilnya.
  	  - `hofog (c)` menghitung h(f(g(c))) dan menampilkan hasilnya.
     
### Cara Kerja Program:
	1. Input: Pengguna diminta untuk memasukkan tiga nilai bilangan bulat a, b, dan c.
	2. Pemrosesan:
  	  - Program pertama-tama menghitung hasil dari fogoh(a), yaitu komposisi fungsi f(g(h(a))):
    	  - Fungsi h(a) dihitung terlebih dahulu, yaitu a + 1.
    	  - Hasilnya kemudian digunakan sebagai input untuk g(x), yaitu g(h(a)) = (a + 1) – 2.
       	- Hasil dari g(h(a)) digunakan sebagai input untuk f(x), yaitu f(g(h(a))) = (a - 1)^2.
   
 	  - Selanjutnya program menghitung `gohof(b)`, yaitu komposisi fungsi g(h(f(b))):
    	  - Fungsi f(b) dihitung terlebih dahulu, yaitu b^2.
   	  - Hasil dari f(b) digunakan sebagai input untuk h(x), yaitu h(f(b)) = b^2 +1.
    	  - Hasil dari h(f(b)) digunakan sebagai input untuk g(x), yaitu g(h(f(b))) = (b^2 + 1) – 2.
     	  - Terakhir program menghitung `hofog(c), yaitu komposisi fungsi h(f(g(c))):
	  - Fungsi g(c) dihitung terlebih dahulu, yaitu  c - 2 .
    	  - Hasil dari g(c) digunakan sebagai input untuk f(x), yaitu f(g(c)) = (c - 2)^2.
       	  - Hasil dari f(g(c)) digunakan sebagai input untuk h(x), yaitu h(f(g(c))) = (c - 2)^2 + 1.
	3. Output: Setelah setiap komposisi fungsi dihitung, hasilnya dicetak ke layar.


### Soal Nomor 3
	Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟.
 	Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut.
	Masukan terdiri dari beberapa tiga baris.
 	Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2,
  	sedangkan baris ketiga adalah koordinat titik sembarang. 
   	Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
  	Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", 
   	"Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
![image](https://github.com/user-attachments/assets/b165597d-c0bb-443c-af66-d79d4e92bef9)
![image](https://github.com/user-attachments/assets/97102008-00ea-4490-b072-2180290930c2)

#### Source Code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

// Fungsi untuk mengecek apakah titik (x, y) berada di dalam lingkaran
func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	// Input dari pengguna
	var cx1, cy1, r1 float64 // Koordinat dan radius lingkaran 1
	var cx2, cy2, r2 float64 // Koordinat dan radius lingkaran 2
	var x, y float64         // Koordinat titik sembarang

	fmt.Println("Masukkan koordinat pusat lingkaran 1 (cx1, cy1) dan radius r1:")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Println("Masukkan koordinat pusat lingkaran 2 (cx2, cy2) dan radius r2:")
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Println("Masukkan koordinat titik sembarang (x, y):")
	fmt.Scan(&x, &y)

	// Cek posisi titik terhadap lingkaran 1 dan lingkaran 2
	diDalamLingkaran1 := diDalam(cx1, cy1, r1, x, y)
	diDalamLingkaran2 := diDalam(cx2, cy2, r2, x, y)

	// Output hasil
	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
#### Output
![image](https://github.com/user-attachments/assets/0559490b-ee1f-4b60-b86d-cda9d18d60c4)
### Deskripsi Program:
	Program ini digunakan untuk menentukan apakah sebuah titik berada di dalam dua lingkaran atau di luar lingkaran.
 	Program akan meminta input dari pengguna berupa koordinat pusat dan radius dari dua lingkaran, serta koordinat sebuah titik sembarang.
  	Selanjutnya, program akan mengecek apakah titik tersebut berada di dalam salah satu atau kedua lingkaran, atau berada di luar keduanya.
### Algoritma Program:
	1. Program mendeklarasikan fungsi jarak(x1, y1, x2, y2) untuk menghitung jarak antara dua titik menggunakan rumus jarak Euclidean
	2. Fungsi diDalam(cx, cy, r, x, y) digunakan untuk menentukan apakah suatu titik (x,y) berada di dalam lingkaran dengan pusat
 	(cx,cy) dan radius r.Jika jarak antara pusat lingkaran dan titik lebih kecil atau sama dengan radius lingkaran, titik dianggap berada di dalam lingkaran.
	3. Dalam fungsi main:
	• Program menerima input dari pengguna berupa:
	• cx1, cy1, r1: koordinat pusat dan radius lingkaran pertama.
	• cx2, cy2, r2: koordinat pusat dan radius lingkaran kedua.
	• x, y koordinat titik sembarang yang ingin diuji.
	4. Program mengecek apakah titik (x,y) berada di dalam lingkaran 1 atau lingkaran 2:
	• Menggunakan fungsi di dalam untuk mengecek apakah titik berada di dalam lingkaran pertama dan kedua.
	5. Berdasarkan hasil pengecekan, program menampilkan salah satu dari empat kemungkinan:
	• "Titik di dalam lingkaran 1 dan 2" jika titik berada di dalam kedua lingkaran.
	• "Titik di dalam lingkaran 1" jika titik hanya berada di dalam lingkaran pertama.
	• "Titik di dalam lingkaran 2" jika titik hanya berada di dalam lingkaran kedua.
	• "Titik di luar lingkaran 1 dan 2" jika titik berada di luar kedua lingkaran

### Cara Kerja Program:
	1. Input:
	• Pengguna memasukkan koordinat pusat dan radius untuk dua lingkaran, serta koordinat titik sembarang yang ingin diperiksa.
	2. Pemrosesan:
	• Program menghitung jarak antara titik sembarang dengan pusat lingkaran pertama menggunakan fungsi jarak,
 	kemudian membandingkan jarak tersebut dengan radius lingkaran pertama untuk menentukan apakah titik berada di dalam lingkaran.
	• Langkah yang sama dilakukan untuk lingkaran kedua.
	3. Output:
	• Program menampilkan hasil sesuai dengan posisi titik terhadap kedua lingkaran, 
 	apakah di dalam kedua lingkaran, di salah satu lingkaran, atau di luar keduanya.













   
