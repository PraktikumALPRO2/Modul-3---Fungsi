
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
![image](https://github.com/user-attachments/assets/8e11926f-6892-4723-9fb9-f75fd32cf60b)


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
![image](https://github.com/user-attachments/assets/2ddb474f-8031-4c4d-9e57-c30635cea247)


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
![image](https://github.com/user-attachments/assets/b6556acc-04a8-4f86-9c03-1c11b75d8c86)


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

3. Fungsi `permutation` didefinisikan untuk menghitung permutasi, di mana fungsi ini akan memanggil fungsi `factorial` untuk menghitung nilai faktorial `n` dan `n-r`, kemudian membagi hasilnya untuk mendapatkan nilai permutasi.

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
![image](https://github.com/user-attachments/assets/03409a5d-c8bb-4f2d-90d0-703e24f4b27f)


#### Deskripsi Program
Program ini ditulis dalam bahasa Go untuk menghitung hasil dari beberapa komposisi fungsi matematika. Terdapat tiga fungsi dasar: `f(x)` yang mengkuadratkan `x`, `g(x)` yang mengurangi `x` dengan dua, dan `h(x)` yang menambahkan satu pada `x`. Program meminta pengguna untuk memasukkan tiga bilangan bulat, kemudian menghitung dan menampilkan hasil dari komposisi fungsi `f(g(h(x)))`, `g(h(f(x)))`, dan `h(f(g(x)))` untuk input tersebut. Hasil dari setiap komposisi ditampilkan secara terpisah.
#### Algoritma Program
1. Program dimulai.

2. Program mendeklarasikan tiga fungsi matematika: fungsi `f(x)` untuk menghitung kuadrat dari x, fungsi `g(x)` untuk mengurangi x dengan 2, dan fungsi `h(x)` untuk menambah x dengan 1.

3. Program mendeklarasikan tiga fungsi komposisi:

- `fogoh(x)` yang menghitung komposisi fungsi `f(g(h(x)))`.
- `gohof(x)` yang menghitung komposisi fungsi `g(h(f(x)))`.
- `hofog(x)` yang menghitung komposisi fungsi `h(f(g(x)))`.

4. Program meminta input dari pengguna untuk tiga bilangan bulat yang akan disimpan dalam variabel `a`, `b`, dan `c`.

5. Program menghitung hasil dari komposisi fungsi `fogoh(a)`, yaitu `f(g(h(a)))`.

6. Program menghitung hasil dari komposisi fungsi `gohof(b)`, yaitu `g(h(f(b)))`.

7. Program menghitung hasil dari komposisi fungsi `hofog(c)`, yaitu `h(f(g(c)))`.

8. Program menampilkan hasil dari `fogoh(a)` pada baris pertama.

9. Program menampilkan hasil dari `gohof(b)` pada baris kedua.

10. Program menampilkan hasil dari `hofog(c)` pada baris ketiga.

11. Program selesai.
#### Cara Kerja Program
1. Program mendeklarasikan tiga fungsi matematika:`f(x)` untuk menghitung kuadrat, `g(x)` untuk mengurangi nilai dengan 2, dan `h(x)` untuk menambah nilai dengan 1.

2. Program mendefinisikan tiga fungsi komposisi, yaitu `fogoh(x)` yang menghitung `f(g(h(x)))`, `gohof(x)` yang menghitung `g(h(f(x)))`, dan `hofog(x)` yang menghitung `h(f(g(x)))`.

3. Program meminta input dari pengguna berupa tiga bilangan bulat yang disimpan dalam variabel `a`, `b`, dan `c`.

4. Program menghitung hasil dari komposisi fungsi `fogoh(a)`, yaitu menerapkan fungsi `h` pada `a`, kemudian menerapkan fungsi `g` pada hasil tersebut, dan terakhir menerapkan fungsi `f` pada hasil dari `g`.

5. Program menghitung hasil dari komposisi fungsi `gohof(b)`, yaitu menerapkan fungsi `f` pada `b`, lalu fungsi `h` pada hasil tersebut, dan terakhir fungsi `g`.

6. Program menghitung hasil dari komposisi fungsi `hofog(c)`, yaitu menerapkan fungsi `g` pada `c`, kemudian fungsi `f` pada hasil tersebut, dan terakhir fungsi `h`.

7. Program menampilkan hasil dari `fogoh(a)` sebagai output pertama.

8. Program menampilkan hasil dari `gohof(b)` sebagai output kedua.

9. Program menampilkan hasil dari `hofog(c)` sebagai output ketiga.

10. Program selesai dijalankan.

### 3. Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
![Cuplikan layar 2024-10-10 135215](https://github.com/user-attachments/assets/9b921d6b-79bf-4dc0-b8b7-0fefb916f31f)
![Cuplikan layar 2024-10-10 135233](https://github.com/user-attachments/assets/47418b31-16c0-4cb1-ad2a-8b3efe995f23)
#### Source Code :
```go
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func diDalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	// Input untuk lingkaran 1
	fmt.Println("Masukkan koordinat pusat (cx, cy) dan radius r lingkaran 1:")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input untuk lingkaran 2
	fmt.Println("Masukkan koordinat pusat (cx, cy) dan radius r lingkaran 2:")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input untuk titik sembarang
	fmt.Println("Masukkan koordinat titik sembarang (x, y):")
	fmt.Scan(&x, &y)

	// Mengecek posisi titik terhadap lingkaran 1 dan lingkaran 2
	dalamLingkaran1 := diDalamLingkaran(cx1, cy1, r1, x, y)
	dalamLingkaran2 := diDalamLingkaran(cx2, cy2, r2, x, y)

	// Menentukan keluaran berdasarkan posisi titik
	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
#### Output
![image](https://github.com/user-attachments/assets/47a2189f-d4fa-4a15-b63b-2b0b7fc432eb)
#### Deskripsi Program
Program ini bertujuan untuk menentukan apakah sebuah titik sembarang berada di dalam atau di luar dua lingkaran yang diberikan. Lingkaran-lingkaran tersebut didefinisikan oleh pengguna melalui koordinat pusat dan radiusnya, sedangkan titik sembarang didefinisikan oleh koordinatnya. Program dimulai dengan mendeklarasikan fungsi `jarak` untuk menghitung jarak antara dua titik berdasarkan rumus Euclidean. Fungsi `diDalamLingkaran` kemudian digunakan untuk memeriksa apakah jarak antara titik sembarang dan pusat lingkaran lebih kecil atau sama dengan radius lingkaran, yang menunjukkan bahwa titik tersebut berada di dalam lingkaran. Setelah pengguna memasukkan data lingkaran pertama, lingkaran kedua, dan titik sembarang, program akan mengevaluasi apakah titik tersebut berada di dalam satu atau kedua lingkaran, atau di luar keduanya. Hasilnya akan ditampilkan dalam bentuk pesan yang menyatakan apakah titik berada di dalam lingkaran 1, lingkaran 2, kedua lingkaran, atau di luar keduanya.
#### Algoritma Program
1. Program dimulai.

2. Program mendeklarasikan fungsi `jarak(x1, y1, x2, y2)` untuk menghitung jarak antara dua titik berdasarkan rumus Euclidean.

3. Program mendeklarasikan fungsi `diDalamLingkaran(cx, cy, r, x, y)` untuk menentukan apakah titik berada di dalam atau di luar lingkaran berdasarkan jarak ke pusat lingkaran.

4. Program meminta input dari pengguna untuk koordinat pusat `(cx1, cy1)` dan radius `r1` dari lingkaran pertama.

5. Program meminta input dari pengguna untuk koordinat pusat `(cx2, cy2)` dan radius `r2` dari lingkaran kedua.

6. Program meminta input dari pengguna untuk koordinat titik sembarang `(x, y)`.

7. Program menggunakan fungsi `diDalamLingkaran` untuk memeriksa apakah titik berada di dalam lingkaran pertama.

8. Program menggunakan fungsi `diDalamLingkaran` untuk memeriksa apakah titik berada di dalam lingkaran kedua.

9. Program mengecek hasil dari kedua pemeriksaan dan menentukan apakah titik berada di dalam lingkaran 1 dan 2, hanya lingkaran 1, hanya lingkaran 2, atau di luar kedua lingkaran.

10. Program menampilkan hasil sesuai dengan kondisi yang ditemukan.

11. Program selesai.
#### Cara Kerja Program
1. Program dimulai dengan mendeklarasikan dua fungsi utama, yaitu `jarak` untuk menghitung jarak antara dua titik, dan `diDalamLingkaran` untuk memeriksa apakah suatu titik berada di dalam lingkaran.

2. Fungsi `jarak(x1, y1, x2, y2)` menghitung jarak Euclidean antara dua titik `(x1, y1)` dan `(x2, y2)` dan mengembalikan nilai jarak sebagai bilangan desimal `(float64)`.

3. Fungsi `diDalamLingkaran(cx, cy, r, x, y)` memeriksa apakah titik `(x, y)` berada di dalam lingkaran yang berpusat di `(cx, cy)` dengan radius `r` menggunakan hasil dari fungsi `jarak`.

4. Program meminta input dari pengguna berupa koordinat pusat `(cx1, cy1)` dan radius `r1` untuk lingkaran pertama.

5. Program meminta input dari pengguna untuk koordinat pusat `(cx2, cy2)` dan radius `r2` untuk lingkaran kedua.

6. Program meminta input dari pengguna untuk koordinat titik sembarang `(x, y)`.

7. Program memanggil fungsi `diDalamLingkaran` untuk memeriksa apakah titik `(x, y)` berada di dalam lingkaran pertama.

8. Program memanggil fungsi `diDalamLingkaran` untuk memeriksa apakah titik `(x, y)` berada di dalam lingkaran kedua.

9. Berdasarkan hasil pemeriksaan, program menentukan apakah titik tersebut berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar kedua lingkaran.

10. Program menampilkan hasil pemeriksaan sesuai kondisi yang ditemukan, apakah titik berada di dalam salah satu atau kedua lingkaran, atau di luar keduanya.
