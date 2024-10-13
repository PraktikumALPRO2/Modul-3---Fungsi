<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL III</strong>
  <br>
  <strong>Fungsi</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
 Aditya Sulistiaawan
  <br>
  2311102193
  <br>
  IF-11-05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>


## <strong> Dasar Teori </strong>

Fungsi adalah salah satu konsep fundamental dalam pemrograman yang memungkinkan programmer untuk mengorganisir dan mengelompokkan serangkaian instruksi menjadi unit yang dapat digunakan kembali. Berikut adalah beberapa aspek penting dari fungsi:

### 1. Definisi:
   Fungsi adalah blok kode yang dirancang untuk melakukan tugas tertentu. Fungsi hanya akan dijalankan ketika dipanggil dan dapat menerima data (parameter), memprosesnya, dan mengembalikan hasil.

### 2. Tujuan Penggunaan Fungsi:
   - Meningkatkan modularitas kode
   - Mengurangi redundansi dengan memungkinkan penggunaan kembali kode
   - Meningkatkan keterbacaan dengan memecah program menjadi bagian-bagian yang lebih kecil dan terkelola
   - Mempermudah pemeliharaan dan debugging kode

### 3. Struktur Umum Fungsi:
   - Nama fungsi: Identifier unik untuk fungsi
   - Parameter: Data yang diterima fungsi (opsional)
   - Body: Serangkaian instruksi yang mendefinisikan operasi fungsi
   - Return value: Nilai yang dikembalikan oleh fungsi (opsional)

### 4. Jenis-jenis Fungsi:
   - Built-in functions: Fungsi yang sudah tersedia dalam bahasa pemrograman
   - User-defined functions: Fungsi yang dibuat oleh programmer
   - Anonymous functions / Lambda functions: Fungsi tanpa nama yang biasanya digunakan untuk operasi singkat

### 5. Konsep Penting dalam Fungsi:
   - Pemanggilan fungsi (Function call)
   - Passing parameter
   - Scope variabel (local dan global)
   - Rekursi: Fungsi yang memanggil dirinya sendiri

### 6. Manfaat Penggunaan Fungsi:
   - Abstraksi: Menyembunyikan kompleksitas implementasi
   - Reusabilitas: Kode dapat digunakan kembali di berbagai bagian program
   - Testabilitas: Memudahkan pengujian unit kode
   - Kolaborasi: Memungkinkan pembagian tugas dalam tim pengembangan

### 7. Best Practices dalam Penulisan Fungsi:
   - Fungsi sebaiknya melakukan satu tugas spesifik (Single Responsibility Principle)
   - Nama fungsi harus deskriptif dan mencerminkan tujuannya
   - Batasi jumlah parameter untuk meningkatkan keterbacaan
   - Dokumentasikan fungsi dengan komentar atau docstrings

Fungsi merupakan komponen penting dalam paradigma pemrograman terstruktur dan berorientasi objek. Pemahaman yang baik tentang fungsi sangat penting untuk mengembangkan kode yang efisien, mudah dipelihara, dan dapat digunakan kembali.


<strong> Deklarasi Fungsi
</strong>

Deklarasi fungsi di Golang menggunakan kata kunci func, diikuti oleh nama fungsi, daftar parameter (opsional), tipe data hasil pengembalian (opsional), dan badan fungsi yang ditulis dalam kurung kurawal {}.

*Contoh*
``` go
func functionName(parameter1 type, parameter2 type, ...) returnType {
    // Badan fungsi
}
```
- functionName: Nama fungsi yang akan digunakan untuk memanggilnya.
- parameter: Nilai yang diterima fungsi sebagai input (opsional).
- returnType: Tipe data hasil pengembalian (opsional).
- Badan fungsi: Kode yang dijalankan saat fungsi dipanggil.

<strong> Cara Memanggil Function</strong>

Setelah dideklarasikan, fungsi dapat dipanggil di dalam kode utama program. Pemanggilan fungsi dilakukan dengan menuliskan nama fungsi diikuti oleh tanda kurung `()`. Jika fungsi memiliki parameter, maka argumen yang sesuai harus diberikan di dalam tanda kurung.

*Contoh*
```go
package main

import "fmt"

// Fungsi dengan parameter
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    // Pemanggilan fungsi dengan argumen
    greet("John")  // Hello, John
    greet("Alice") // Hello, Alice
}
```


## <strong> Guided </strong>

### <h2>GUIDED 1</h2>

#### Source Code

```go
package main 

import "fmt"
 
func main(){
	var a,b int 
	fmt.Scan(&a, &b) 
	if a >= b {
		fmt.Println(permutasi(a,b))
	}else{
		fmt.Println(permutasi(b,a))
	}
}
func faktorial(n int) int{ 
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ { 
		hasil = hasil * i
	}

	return hasil
}
func permutasi(n,r int) int {
	return faktorial(n) / faktorial(n-r)
}
```

#### Screenshot Source Code
![SC](https://github.com/user-attachments/assets/e977240d-e959-481f-b9a8-77818f00196c)


#### Output
![Gui1](https://github.com/user-attachments/assets/51f1bee9-d0b0-402a-9aca-fbe34751e7be)


#### Deskripsi Program : 
Program ini ditulis dalam bahasa Go (Golang) dan berfungsi untuk menghitung permutasi dari dua angka yang diinputkan oleh pengguna.

#### Algoritma dan Cara Kerja Program

1. Mulai program
2. Deklarasikan variabel a dan b sebagai integer
3. Minta input dari pengguna untuk nilai a dan b
4. Jika a lebih besar atau sama dengan b, maka:
   - Panggil fungsi permutasi(a, b)
   - Tampilkan hasil permutasi
5. Jika tidak (a lebih kecil dari b), maka:
   - Panggil fungsi permutasi(b, a)
   - Tampilkan hasil permutasi
6. Definisikan fungsi faktorial(n):
   a. Inisialisasi variabel hasil = 1
   b. Untuk i dari 1 sampai n, lakukan:
      - hasil = hasil * i
   c. Kembalikan nilai hasil
7. Definisikan fungsi permutasi(n, r):
   a. Hitung faktorial(n)
   b. Hitung faktorial(n-r)
   c. Hitung hasil = faktorial(n) / faktorial(n-r)
   d. Kembalikan nilai hasil
8. Selesai


#### Cara Kerja*

1. Program meminta pengguna memasukkan dua angka.
2. Program menentukan angka mana yang lebih besar.
3. Fungsi `permutasi()` dipanggil dengan angka yang lebih besar sebagai `n` dan yang lebih kecil sebagai `r`.
4. Fungsi `permutasi()` memanggil fungsi `faktorial()` untuk menghitung faktorial yang diperlukan.
5. Hasil perhitungan permutasi dicetak ke layar.

Program ini mendemonstrasikan penggunaan fungsi, pengambilan input dari pengguna, pengambilan keputusan dengan `if-else`, dan perhitungan matematika dasar dalam pemrograman Go.

### <h2>GUIDED 2</h2>

#### Source Code

```go
package main

import "fmt"

func hitungLuas(sisi float64) float64 {
	return sisi * sisi
}

func hitungKeliling(sisi float64) float64 {
	return 4 * sisi
}

func main() {
	var sisi float64

	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi)

	luas := hitungLuas(sisi)
	keliling := hitungKeliling(sisi)

	fmt.Printf("Luas persegi adalah: %.2f\n", luas)
	fmt.Printf("Keliling persegi adalah: %.2f\n", keliling)
}
```

#### Source Code
![SC](https://github.com/user-attachments/assets/65d918b5-dc30-48a3-b647-9983985d9245)


#### Output
![Gui2](https://github.com/user-attachments/assets/d61e062a-ebaa-435e-b3a8-25ec5a2f3b1e)



#### Deskripsi Program : 
Program ini ditulis dalam bahasa pemrograman Go (Golang) dan berfungsi untuk menghitung luas dan keliling dari sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna. Program ini menggunakan dua fungsi terpisah: satu untuk menghitung luas dan satu lagi untuk menghitung keliling. Setelah pengguna memasukkan nilai sisi, program akan menampilkan hasil perhitungan tersebut dengan format yang rapi.

#### Algoritma 
1. **Inisialisasi**:
   - Import paket `fmt` untuk input/output.
   
2. **Definisi Fungsi**:
   - `hitungLuas(sisi float64)`: Menghitung luas persegi dengan rumus $$ \text{luas} = \text{sisi} \times \text{sisi} $$.
   - `hitungKeliling(sisi float64)`: Menghitung keliling persegi dengan rumus $$ \text{keliling} = 4 \times \text{sisi} $$.

3. **Fungsi `main`**:
   - Deklarasi variabel `sisi` bertipe `float64`.
   - Mencetak pesan untuk meminta pengguna memasukkan panjang sisi persegi.
   - Menggunakan `fmt.Scanln(&sisi)` untuk membaca input dari pengguna.
   - Memanggil fungsi `hitungLuas` dan menyimpan hasilnya dalam variabel `luas`.
   - Memanggil fungsi `hitungKeliling` dan menyimpan hasilnya dalam variabel `keliling`.
   - Mencetak hasil luas dan keliling dengan format dua desimal.



#### Cara Kerja


1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan panjang sisi persegi melalui konsol.

2. **Proses Perhitungan**:
   - Setelah menerima input, program memanggil fungsi `hitungLuas` untuk menghitung luas berdasarkan sisi yang diberikan.
   - Selanjutnya, program memanggil fungsi `hitungKeliling` untuk menghitung keliling dari persegi tersebut.

3. **Output Hasil**:
   - Program menampilkan hasil perhitungan luas dan keliling ke layar dengan format yang mudah dibaca, menggunakan dua angka di belakang koma.
     
## <strong> Unguided </strong>

### <h2> UNGUIDED 1 </h2>

#### Study Case
Minggu Ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kullah matematika diskett untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, Iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kallan membantu Jonas? (tidak tentunya ya :p)

#### Source Code

```go
// Aditya Sulistiawan
// 2311102193
// 11-05

package main

import "fmt"

func main() {
	var a, b, c, d int
5
	// fungsi untuk menginput bilangan a, b, c, dan d
	fmt.Print("Bilangan a: ")
	fmt.Scan(&a)

	fmt.Print("Bilangan b: ")
	fmt.Scan(&b)

	fmt.Print("Bilangan c: ")
	fmt.Scan(&c)

	fmt.Print("Bilangan d: ")
	fmt.Scan(&d)

	// Untuk menjelaskan syarat bahwa a>= c dan b>=d
	if a >= c && b >= d {
		// Untuk menampilkan hasil permutasi dan kombinasi dari bilangan a dan c
		fmt.Printf("\nPermutasi (a, c): %d\n", permutasi(a, c))
		fmt.Printf("Kombinasi (a, c): %d\n", kombinasi(a, c))

		// Untuk menampilkan hasil permutasi dan kombinasi dari bilangan b dan d
		fmt.Printf("\nPermutasi (b, d): %d\n", permutasi(b, d))
		fmt.Printf("Kombinasi (b, d): %d\n", kombinasi(b, d))
	} else {
		// Untuk yang tidak memenuhi syarat
		fmt.Println("Input tidak sesuai dengan syarat yang ada")
	}
}

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
```

#### Source Code
![SC](https://github.com/user-attachments/assets/31374b8d-770c-4190-bfa5-231d1488d8ae)


#### Output
![Ungui1](https://github.com/user-attachments/assets/d202d887-911b-4643-adf0-b761ef7962db)



#### Deskripsi Program : 
Program ini ditulis dalam bahasa pemrograman Go (Golang) dan bertujuan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan bulat, yaitu `a` dan `c`, serta `b` dan `d`. Program ini juga memeriksa syarat bahwa `a` harus lebih besar atau sama dengan `c`, dan `b` harus lebih besar atau sama dengan `d`. Jika syarat tersebut tidak terpenuhi, program akan memberikan pesan kesalahan

#### Algoritma dan Cara Kerja Program
1. **Inisialisasi**:
   - Import paket `fmt` untuk input/output.
   
2. **Deklarasi Variabel**:
   - Deklarasikan empat variabel bertipe `int`: `a`, `b`, `c`, dan `d`.

3. **Input Pengguna**:
   - Minta pengguna untuk memasukkan nilai untuk bilangan `a`, `b`, `c`, dan `d`.

4. **Pemeriksaan Syarat**:
   - Periksa apakah `a >= c` dan `b >= d`.
     - Jika syarat terpenuhi:
       - Hitung dan tampilkan permutasi dan kombinasi untuk pasangan `(a, c)` dan `(b, d)`.
     - Jika syarat tidak terpenuhi:
       - Tampilkan pesan kesalahan.

5. **Fungsi Pendukung**:
   - **Fungsi `faktorial(n int) int`**: Menghitung faktorial dari bilangan `n`.
   - **Fungsi `permutasi(n, r int) int`**: Menghitung permutasi menggunakan rumus $$ P(n, r) = \frac{n!}{(n-r)!} $$.
   - **Fungsi `kombinasi(n, r int) int`**: Menghitung kombinasi menggunakan rumus $$ C(n, r) = \frac{n!}{r!(n-r)!} $$.


#### Cara Kerja
1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan empat bilangan bulat: `a`, `b`, `c`, dan `d`.

2. **Pemeriksaan Syarat**:
   - Setelah menerima input, program memeriksa apakah nilai-nilai yang dimasukkan memenuhi syarat (`a >= c` dan `b >= d`).
   - Jika syarat tidak terpenuhi, program akan mencetak pesan kesalahan.

3. **Perhitungan Permutasi dan Kombinasi**:
   - Jika syarat terpenuhi, program akan menghitung permutasi dan kombinasi untuk kedua pasangan bilangan:
     - Untuk pasangan `(a, c)`:
       - Memanggil fungsi `permutasi(a, c)` untuk menghitung permutasi.
       - Memanggil fungsi `kombinasi(a, c)` untuk menghitung kombinasi.
     - Untuk pasangan `(b, d)`:
       - Memanggil fungsi `permutasi(b, d)` untuk menghitung permutasi.
       - Memanggil fungsi `kombinasi(b, d)` untuk menghitung kombinasi.

4. **Output Hasil**:
   - Program menampilkan hasil perhitungan permutasi dan kombinasi ke layar dengan format yang jelas.

Dengan struktur yang terorganisir dengan baik, program ini menyediakan cara yang efektif untuk melakukan perhitungan kombinatorial dasar dengan memanfaatkan fungsi-fungsi pendukung untuk faktorial, permutasi, dan kombinasi.

### <h2> UNGUIDED 2 </h2>

#### Study Case
Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥2 , 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 + 1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔Gℎ(𝑥)H). Tuliskan 𝑓(𝑥), 𝑔(𝑥) dan ℎ(𝑥) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat 𝑎, 𝑏 dan 𝑐 yang dipisahkan oleh spasi.

Keluaran terdiri dari tiga baris. Baris pertama adalah (𝑓𝑜𝑔𝑜ℎ)(𝑎), baris kedua (𝑔𝑜ℎ𝑜𝑓)(𝑏), dan baris ketiga adalah (ℎ𝑜𝑓𝑜𝑔)(𝑐)

#### Source Code

```go
// Aditya Sulistiawan
// 2311102193
// 11-05

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

// Fungsi komposisi f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi komposisi g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi komposisi h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	// Input bilangan a, b, c
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// Menghitung hasil komposisi
	result1 := fogoh(a)
	result2 := gohof(b)
	result3 := hofog(c)

	// Output hasil komposisi
	fmt.Println(result1) // (f o g o h)(a)
	fmt.Println(result2) // (g o h o f)(b)
	fmt.Println(result3) // (h o f o g)(c)
}
```

#### Source Code
![SC](https://github.com/user-attachments/assets/91254b4f-7e72-4f5e-a9b0-a47b70051d61)


#### Output
![Ungui2](https://github.com/user-attachments/assets/399bf6d3-7b7a-4001-957b-1fb3b04156a3)


#### Deskripsi Program : 
Program ini ditulis dalam bahasa pemrograman Go (Golang) dan berfungsi untuk melakukan komposisi fungsi matematika. Terdapat tiga fungsi dasar yang didefinisikan: `f(x)`, `g(x)`, dan `h(x)`, masing-masing dengan rumus tertentu. Program ini kemudian menghitung hasil dari tiga komposisi fungsi yang berbeda menggunakan input dari pengguna. Hasil dari setiap komposisi akan ditampilkan di konsol.

#### Algoritma
1. **Inisialisasi**:
   - Import paket `fmt` untuk input/output.

2. **Definisi Fungsi**:
   - `f(x int) int`: Menghitung nilai $$ x^2 $$.
   - `g(x int) int`: Menghitung nilai $$ x - 2 $$.
   - `h(x int) int`: Menghitung nilai $$ x + 1 $$.
   - `fogoh(x int) int`: Menghitung komposisi $$ f(g(h(x))) $$.
   - `gohof(x int) int`: Menghitung komposisi $$ g(h(f(x))) $$.
   - `hofog(x int) int`: Menghitung komposisi $$ h(f(g(x))) $$.

3. **Fungsi `main`**:
   - Deklarasi variabel `a`, `b`, dan `c` untuk menyimpan input pengguna.
   - Membaca nilai dari pengguna untuk variabel `a`, `b`, dan `c`.
   - Menghitung hasil dari setiap komposisi fungsi.
   - Mencetak hasil ke layar.

#### Cara Kerja
1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan tiga bilangan bulat: `a`, `b`, dan `c`.

2. **Perhitungan Komposisi Fungsi**:
   - Setelah menerima input, program menghitung hasil dari tiga komposisi fungsi:
     - **Komposisi Pertama**: 
       - Memanggil fungsi `fogoh(a)` untuk menghitung $$ f(g(h(a))) $$.
     - **Komposisi Kedua**: 
       - Memanggil fungsi `gohof(b)` untuk menghitung $$ g(h(f(b))) $$.
     - **Komposisi Ketiga**: 
       - Memanggil fungsi `hofog(c)` untuk menghitung $$ h(f(g(c))) $$.

3. **Output Hasil**:
   - Program menampilkan hasil dari setiap komposisi ke layar dengan format yang jelas.

### Contoh Eksekusi

Jika pengguna memasukkan nilai:
-  a = 3 
-  b = 5 
-  c = 1 

Program akan menghitung:
-  f(g(h(3))) = f(g(4)) = f(2) = 4 
-  g(h(f(5))) = g(h(25)) = g(26) = 24 
-  h(f(g(1))) = h(f(-1)) = h(1) = 2 

Hasil yang ditampilkan akan menjadi:
```
4
24
2
```

Dengan cara ini, program memberikan cara yang efisien dan terstruktur untuk mengeksplorasi komposisi fungsi matematika menggunakan input dinamis dari pengguna.

### <h2> UNGUIDED 3 </h2>

#### Study Case
Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut.

Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".


#### Source Code

```go
// Aditya Sulistiawan
// 2311102193
// 11-05

package main

import "fmt"

func hitungKuadratJarak(x1, y1, x2, y2 int) int {
	return (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
}

func dalamLingkaran(cx, cy, r, x, y int) bool {
	return hitungKuadratJarak(cx, cy, x, y) <= r*r
}

func main() {

	var cx1, cy1, r1 int
	fmt.Print("Koordinat pusat dan radius pada lingkaran 1 ")
	fmt.Println()
	fmt.Print("Koordinat pusat x: ")
	fmt.Scan(&cx1)
	fmt.Print("Koordinat pusat y: ")
	fmt.Scan(&cy1)
	fmt.Print("Radius lingkaran 1: ")
	fmt.Scan(&r1)

	var cx2, cy2, r2 int
	fmt.Print("Koordinat pusat dan radius pada lingkaran 2 ")
	fmt.Println()
	fmt.Print("Koordinat pusat x: ")
	fmt.Scan(&cx2)
	fmt.Print("Koordinat pusat y: ")
	fmt.Scan(&cy2)
	fmt.Print("Radius lingkaran 2: ")
	fmt.Scan(&r2)

	var x, y int
	fmt.Print("Koordinat titik sembarang (x, y)")
	fmt.Println()
	fmt.Print("titik x: ")
	fmt.Scan(&x)
	fmt.Print("titik y: ")
	fmt.Scan(&y)

	diLingkaran1 := dalamLingkaran(cx1, cy1, r1, x, y)
	diLingkaran2 := dalamLingkaran(cx2, cy2, r2, x, y)

	fmt.Print("\nPosisi titik: ")
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik tidak berada di dalam lingkaran 1 dan 2")
	}
}
```

#### Source Code
![SC](https://github.com/user-attachments/assets/64951558-f8c8-43d5-94f9-ff7933ddf31c)




#### Output
![Ungui3-1](https://github.com/user-attachments/assets/3d097d99-a049-4b95-ab89-8de87b30239b)
![Ungui3-2](https://github.com/user-attachments/assets/f9f5c139-028c-4307-bcb8-c812b3e58ea9)




#### Deskripsi Program : 
Program ini ditulis dalam bahasa pemrograman Go (Golang) dan bertujuan untuk menentukan apakah sebuah titik berada di dalam dua lingkaran yang berbeda. Program ini meminta pengguna untuk memasukkan koordinat pusat dan radius dari dua lingkaran, serta koordinat titik sembarang. Berdasarkan input tersebut, program akan memeriksa posisi titik relatif terhadap kedua lingkaran dan menampilkan hasilnya.

#### Algoritma dan Cara Kerja Program
1. **Inisialisasi**:
   - Import paket `fmt` untuk input/output.

2. **Definisi Fungsi**:
   - `hitungKuadratJarak(x1, y1, x2, y2 int) int`: Menghitung kuadrat jarak antara dua titik `(x1, y1)` dan `(x2, y2)` dengan rumus:
     
     \text{jarak}^2 = (x2 - x1)^2 + (y2 - y1)^2
     
   - `dalamLingkaran(cx, cy, r, x, y int) bool`: Memeriksa apakah titik `(x, y)` berada dalam lingkaran dengan pusat `(cx, cy)` dan radius `r`. Menggunakan fungsi `hitungKuadratJarak` untuk membandingkan jarak kuadrat dengan kuadrat radius.

3. **Fungsi `main`**:
   - Deklarasi variabel untuk pusat dan radius dari dua lingkaran (`cx1`, `cy1`, `r1`, `cx2`, `cy2`, `r2`).
   - Membaca input dari pengguna untuk pusat dan radius kedua lingkaran.
   - Deklarasi variabel untuk koordinat titik sembarang (`x`, `y`).
   - Membaca input dari pengguna untuk koordinat titik.
   - Memanggil fungsi `dalamLingkaran` untuk memeriksa apakah titik berada dalam masing-masing lingkaran.
   - Menampilkan hasil posisi titik berdasarkan pemeriksaan.

#### Cara Kerja
1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan informasi berikut:
     - Koordinat pusat dan radius dari lingkaran pertama (`cx1`, `cy1`, `r1`).
     - Koordinat pusat dan radius dari lingkaran kedua (`cx2`, `cy2`, `r2`).
     - Koordinat titik sembarang (`x`, `y`).

2. **Perhitungan Jarak**:
   - Program menggunakan fungsi `hitungKuadratJarak` untuk menghitung jarak kuadrat antara pusat lingkaran dan titik yang diberikan.

3. **Pemeriksaan Posisi Titik**:
   - Fungsi `dalamLingkaran` memeriksa apakah jarak kuadrat dari pusat lingkaran ke titik lebih kecil atau sama dengan kuadrat radius. Ini dilakukan untuk kedua lingkaran.

4. **Output Hasil**:
   - Program mencetak posisi titik berdasarkan hasil pemeriksaan:
     - Jika titik berada di dalam kedua lingkaran.
     - Jika titik hanya berada di dalam salah satu lingkaran.
     - Jika titik tidak berada di dalam kedua lingkaran.

## Referensi
[1] Noval Agung Prayogo, *Dasar Pemrograman Golang*. 2024
[2] https://dasarpemrogramangolang.novalagung.com/A-hello-world.html
